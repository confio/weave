package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/cmd/bnsd/x/account"
)

func TestHexbytes(t *testing.T) {
	a := hexbytes("a hexbyte value")
	raw, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("cannot marshal: %s", err)
	}
	var b hexbytes
	if err := json.Unmarshal(raw, &b); err != nil {
		t.Fatalf("cannot unmarshal: %s", err)
	}
	if !bytes.Equal(a, b) {
		t.Fatalf("%q != %q", a, b)
	}
}

func TestBnsClientMock(t *testing.T) {
	// Just to be sure, test the mock.

	result := abciQueryResponse{
		Response: abciQueryResponseResponse{
			Key:   []byte("foo"),
			Value: []byte("bar"),
		},
	}
	bns := bnsClientMock{Results: map[string]abciQueryResponse{
		"/foo": result,
	}}
	var response abciQueryResponse
	if err := bns.Get(context.Background(), "/foo", &response); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(response, result) {
		t.Fatalf("unexpected response: %+v", response)
	}
}

type bnsClientMock struct {
	Results map[string]abciQueryResponse
	Err     error
}

func (mock *bnsClientMock) Get(ctx context.Context, path string, dest interface{}) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	resp, ok := mock.Results[path]
	if !ok {
		raw, _ := url.PathUnescape(path)
		return fmt.Errorf("no result declared in mock for %q (%q)", path, raw)
	}

	v := reflect.ValueOf(dest)
	// Below panics if cannot be fullfilled. User did something wrong and
	// this is test so panic is acceptable.
	src := reflect.ValueOf(resp)
	v.Elem().Set(src)

	return mock.Err
}

func TestAccountAccountssHandler(t *testing.T) {
	bns := &bnsClientMock{
		Results: map[string]abciQueryResponse{
			"/abci_query?data=%22%3A%22&path=%22%2Faccounts%3Frange%22": newAbciQueryResponse(t,
				[][]byte{
					[]byte("first"),
					[]byte("second"),
				},
				[]weave.Persistent{
					&account.Account{Name: "first", Domain: "adomain"},
					&account.Account{Name: "second", Domain: "adomain"},
				}),
		},
	}
	h := AccountAccountsHandler{bns: bns}

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	assertAPIResponse(t, w, []KeyValue{
		{
			Key:   []byte("first"),
			Value: &account.Account{Name: "first", Domain: "adomain"},
		},
		{
			Key:   []byte("second"),
			Value: &account.Account{Name: "second", Domain: "adomain"},
		},
	})
}

func TestAccountAccountssHandlerOffsetAndFilter(t *testing.T) {
	bns := &bnsClientMock{
		Results: map[string]abciQueryResponse{
			"/abci_query?data=%2261646f6d61696e%3A6669727374%3A%22&path=%22%2Faccounts%2Fdomain%3Frange%22": newAbciQueryResponse(t, nil, nil),
		},
	}
	h := AccountAccountsHandler{bns: bns}

	r, _ := http.NewRequest("GET", "/?offset=6669727374&domain=adomain", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	assertAPIResponse(t, w, nil)
}

func TestAccountDomainsHandler(t *testing.T) {
	bns := &bnsClientMock{
		Results: map[string]abciQueryResponse{
			"/abci_query?data=%22%3A%22&path=%22%2Fdomains%3Frange%22": newAbciQueryResponse(t,
				[][]byte{
					[]byte("first"),
					[]byte("second"),
				},
				[]weave.Persistent{
					&account.Domain{Domain: "f"},
					&account.Domain{Domain: "s"},
				}),
			"/abci_query?data=%227365636f6e64%3A%22&path=%22%2Fdomains%3Frange%22": newAbciQueryResponse(t, nil, nil),
		},
	}
	h := AccountDomainsHandler{bns: bns}

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	assertAPIResponse(t, w, []KeyValue{
		{
			Key:   []byte("first"),
			Value: &account.Domain{Domain: "f"},
		},
		{
			Key:   []byte("second"),
			Value: &account.Domain{Domain: "s"},
		},
	})
}

func TestAccountDomainsHandlerFilter(t *testing.T) {
	bns := &bnsClientMock{
		Results: map[string]abciQueryResponse{
			"/abci_query?data=%22%3A%22&path=%22%2Fdomains%3Frange%22": newAbciQueryResponse(t,
				[][]byte{
					[]byte("first"),
					[]byte("second"),
					[]byte("third"),
				},
				[]weave.Persistent{
					&account.Domain{Domain: "1", Admin: []byte("a1")},
					&account.Domain{Domain: "2", Admin: []byte("a2")},
					&account.Domain{Domain: "3", Admin: []byte("a1")},
				},
			),
			"/abci_query?data=%227468697264%3A%22&path=%22%2Fdomains%3Frange%22": newAbciQueryResponse(t,
				[][]byte{
					[]byte("third"),
					[]byte("fourth"),
					[]byte("fifth"),
				},
				[]weave.Persistent{
					&account.Domain{Domain: "3", Admin: []byte("a1")},
					&account.Domain{Domain: "4", Admin: []byte("a1")},
					&account.Domain{Domain: "5", Admin: []byte("a3")},
				},
			),
			"/abci_query?data=%226669667468%3A%22&path=%22%2Fdomains%3Frange%22": newAbciQueryResponse(t, nil, nil),
		},
	}
	h := AccountDomainsHandler{bns: bns}

	r, _ := http.NewRequest("GET", "/?admin=6131", nil) // admin=a1
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)

	assertAPIResponse(t, w, []KeyValue{
		{
			Key:   []byte("first"),
			Value: &account.Domain{Domain: "1", Admin: []byte("a1")},
		},
		{
			Key:   []byte("third"),
			Value: &account.Domain{Domain: "3", Admin: []byte("a1")},
		},
		{
			Key:   []byte("fourth"),
			Value: &account.Domain{Domain: "4", Admin: []byte("a1")},
		},
	})
}

func newAbciQueryResponse(t testing.TB, keys [][]byte, models []weave.Persistent) abciQueryResponse {
	t.Helper()
	k, v := serializePairs(t, keys, models)

	return abciQueryResponse{
		Response: abciQueryResponseResponse{
			Key:   k,
			Value: v,
		},
	}
}

func assertAPIResponse(t testing.TB, w *httptest.ResponseRecorder, want []KeyValue) {
	t.Helper()

	if w.Code != http.StatusOK {
		t.Log(w.Body)
		t.Fatalf("response code %d", w.Code)
	}

	var payload struct {
		Objects json.RawMessage
	}
	if err := json.NewDecoder(w.Body).Decode(&payload); err != nil {
		t.Fatalf("cannot decode JSON serialized body: %s", err)
	}

	// We cannot unmarshal returned JSON because KeyValue structure does
	// not declare what type Value is. Instead of comparing Go objects,
	// compare JSON output. We know what is the expected JSON content for
	// given KeyValue collection.
	rawGot := []byte(payload.Objects)

	rawWant, err := json.MarshalIndent(want, "", "\t")
	if err != nil {
		t.Fatalf("cannot JSON serialize expected result: %s", err)
	}

	// Because rawGot is part of a bigger JSON message its indentation
	// differs. Indentation is not relevant so it can be removed for
	// comparison.
	if !bytes.Equal(removeTabs(rawGot), removeTabs(rawWant)) {
		t.Logf("want JSON response:\n%s", string(rawWant))
		t.Logf("got JSON response:\n%s", string(rawGot))
		t.Fatal("unexpected response")
	}
}

func removeTabs(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("\t"), []byte(""))
}
