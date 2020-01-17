package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/iov-one/weave/cmd/bnsd/x/account"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

type InfoHandler struct{}

func (h *InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	JSONResp(w, http.StatusOK, struct {
		BuildHash    string `json:"build_hash"`
		BuildVersion string `json:"build_version"`
	}{
		BuildHash:    buildHash,
		BuildVersion: buildVersion,
	})
}

type BlocksHandler struct {
	bns BnsClient
}

func (h *BlocksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	heightStr := lastChunk(r.URL.Path)
	if heightStr == "" {
		JSONRedirect(w, http.StatusSeeOther, "/blocks/1")
		return
	}
	height, err := strconv.ParseInt(heightStr, 10, 64)
	if err != nil {
		JSONErr(w, http.StatusNotFound, "block height must be a number")
		return
	}

	// We do not care about payload, proxy all!
	var payload json.RawMessage
	if err := h.bns.Get(r.Context(), fmt.Sprintf("/block?height=%d", height), &payload); err != nil {
		log.Printf("bns block height info: %s", err)
		JSONErr(w, http.StatusBadGateway, http.StatusText(http.StatusBadGateway))
		return
	}
	JSONResp(w, http.StatusOK, payload)
}

// lastChunk returns last path chunk - everything after the last `/` character.
// For example LAST in /foo/bar/LAST and empty string in /foo/bar/
func lastChunk(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			return path[i+1:]
		}
	}
	return path
}

// DefaultHandler is used to handle the request that no other handler wants.
type DefaultHandler struct{}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// No trailing slash.
	if len(r.URL.Path) > 1 && r.URL.Path[len(r.URL.Path)-1] == '/' {
		path := strings.TrimRight(r.URL.Path, "/")
		JSONRedirect(w, http.StatusPermanentRedirect, path)
		return
	}
	JSONErr(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
}

type AccountDomainsHandler struct {
	bns BnsClient
}

func (h *AccountDomainsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var it ABCIIterator
	q := r.URL.Query()
	offset := extractIDFromKey(q.Get("offset"))
	if admin := q.Get("admin"); len(admin) > 0 {
		it = ABCIRangeQuery(r.Context(), h.bns, "/domains/admin", fmt.Sprintf("%s:%x:", admin, offset))
	} else {
		it = ABCIRangeQuery(r.Context(), h.bns, "/domains", fmt.Sprintf("%x:", offset))
	}

	var objects []KeyValue
fetchDomains:
	for {
		var model account.Domain
		switch key, err := it.Next(&model); {
		case err == nil:
			objects = append(objects, KeyValue{
				Key:   key,
				Value: &model,
			})
			if len(objects) == paginationMaxItems {
				break fetchDomains
			}
		case errors.ErrIteratorDone.Is(err):
			break fetchDomains
		default:
			log.Printf("domain ABCI query: %s", err)
			JSONErr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}
	JSONResp(w, http.StatusOK, struct {
		Objects []KeyValue `json:"objects"`
	}{
		Objects: objects,
	})
}

type AccountAccountDetailHandler struct {
	bns BnsClient
}

func (h *AccountAccountDetailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accountKey := lastChunk(r.URL.Path)
	var acc account.Account
	switch err := ABCIKeyQuery(r.Context(), h.bns, "/accounts", []byte(accountKey), &acc); {
	case err == nil:
		JSONResp(w, http.StatusOK, acc)
	case errors.ErrNotFound.Is(err):
		JSONErr(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
	default:
		log.Printf("account ABCI query: %s", err)
		JSONErr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

type AccountAccountsHandler struct {
	bns BnsClient
}

func (h *AccountAccountsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if !atMostOne(q, "domain", "owner") {
		JSONErr(w, http.StatusBadRequest, "At most one filter can be used at a time.")
		return
	}

	var it ABCIIterator
	offset := extractIDFromKey(q.Get("offset"))
	if d := q.Get("domain"); len(d) > 0 {
		it = ABCIRangeQuery(r.Context(), h.bns, "/accounts/domain", fmt.Sprintf("%x:%x:", d, offset))
	} else if o := q.Get("owner"); len(o) > 0 {
		it = ABCIRangeQuery(r.Context(), h.bns, "/accounts/owner", fmt.Sprintf("%s:%x:", o, offset))
	} else {
		it = ABCIRangeQuery(r.Context(), h.bns, "/accounts", fmt.Sprintf("%x:", offset))
	}

	var objects []KeyValue
fetchAccounts:
	for {
		var acc account.Account
		switch key, err := it.Next(&acc); {
		case err == nil:
			objects = append(objects, KeyValue{
				Key:   key,
				Value: &acc,
			})
			if len(objects) == paginationMaxItems {
				break fetchAccounts
			}
		case errors.ErrIteratorDone.Is(err):
			break fetchAccounts
		default:
			log.Printf("account ABCI query: %s", err)
			JSONErr(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
			return
		}
	}

	JSONResp(w, http.StatusOK, struct {
		Objects []KeyValue `json:"objects"`
	}{
		Objects: objects,
	})
}

// atMostOne returns true if at most one non empty value from given list of
// names exists in the query.
func atMostOne(query url.Values, names ...string) bool {
	var nonempty int
	for _, name := range names {
		if query.Get(name) != "" {
			nonempty++
		}
		if nonempty > 1 {
			return false
		}
	}
	return true
}

func extractIDFromKey(key string) []byte {
	raw, err := hex.DecodeString(key)
	if err != nil {
		// Cannot decode, return everything.
		return []byte(key)
	}
	for i, c := range raw {
		if c == ':' {
			return raw[i+1:]
		}
	}
	return raw
}

// paginationMaxItems defines how many items should a single result return.
// This values should not be greater than orm.queryRangeLimit so that each
// query returns enough results.
const paginationMaxItems = 50

type KeyValue struct {
	Key   hexbytes  `json:"key"`
	Value orm.Model `json:"value"`
}

// hexbytes is a byte type that JSON serialize to hex encoded string.
type hexbytes []byte

func (b hexbytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(b))
}

func (b *hexbytes) UnmarshalJSON(enc []byte) error {
	var s string
	if err := json.Unmarshal(enc, &s); err != nil {
		return err
	}
	val, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	*b = val
	return nil
}

// JSONResp write content as JSON encoded response.
func JSONResp(w http.ResponseWriter, code int, content interface{}) {
	b, err := json.MarshalIndent(content, "", "\t")
	if err != nil {
		log.Printf("cannot JSON serialize response: %s", err)
		code = http.StatusInternalServerError
		b = []byte(`{"errors":["Internal Server Errror"]}`)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)

	const MB = 1 << (10 * 2)
	if len(b) > MB {
		log.Printf("response JSON body is huge: %d", len(b))
	}
	_, _ = w.Write(b)
}

// JSONErr write single error as JSON encoded response.
func JSONErr(w http.ResponseWriter, code int, errText string) {
	JSONErrs(w, code, []string{errText})
}

// JSONErrs write multiple errors as JSON encoded response.
func JSONErrs(w http.ResponseWriter, code int, errs []string) {
	resp := struct {
		Errors []string `json:"errors"`
	}{
		Errors: errs,
	}
	JSONResp(w, code, resp)
}

// JSONRedirect return redirect response, but with JSON formatted body.
func JSONRedirect(w http.ResponseWriter, code int, urlStr string) {
	w.Header().Set("Location", urlStr)
	var content = struct {
		Code     int
		Location string
	}{
		Code:     code,
		Location: urlStr,
	}
	JSONResp(w, code, content)
}
