package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/iov-one/weave"
	weaveapp "github.com/iov-one/weave/app"
	"github.com/iov-one/weave/cmd/bnsd/app"
	"github.com/iov-one/weave/x/username"
)

func cmdRegisterUsername(input io.Reader, output io.Writer, args []string) error {
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Create a transaction for registering a username.
		`)
		fl.PrintDefaults()
	}
	var (
		nameFl       = fl.String("name", "", "Name part of the username. For example 'alice'")
		namespaceFl  = fl.String("ns", "iov", "Namespace (domain) part of the username. For example 'iov'")
		blockchainFl = fl.String("bc", "", "Blockchain network ID.")
		addressFl    = flHex(fl, "addr", "", "Hex encoded blochain address on this network.")
	)
	fl.Parse(args)

	uname, err := username.ParseUsername(*nameFl + "*" + *namespaceFl)
	if err != nil {
		return fmt.Errorf("given data produce an invalid username: %s", err)
	}

	msg := username.RegisterTokenMsg{
		Metadata: &weave.Metadata{Schema: 1},
		Username: uname,
		Target: username.Location{
			BlockchainID: *blockchainFl,
			Address:      *addressFl,
		},
	}
	if err := msg.Validate(); err != nil {
		return fmt.Errorf("given data produce an invalid message: %s", err)
	}

	tx := &app.Tx{
		Sum: &app.Tx_RegisterTokenMsg{
			RegisterTokenMsg: &msg,
		},
	}
	_, err = writeTx(output, tx)
	return err
}

func cmdResolveUsername(input io.Reader, output io.Writer, args []string) error {
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Create a transaction for registering a username.
		`)
		fl.PrintDefaults()
	}
	var (
		tmAddrFl = fl.String("tm", env("BNSCLI_TM_ADDR", "https://bns.NETWORK.iov.one:443"),
			"Tendermint node address. Use proper NETWORK name. You can use BNSCLI_TM_ADDR environment variable to set it.")
		nameFl      = fl.String("name", "", "Name part of the username. For example 'alice'")
		namespaceFl = fl.String("ns", "iov", "Namespace (domain) part of the username. For example 'iov'")
	)
	fl.Parse(args)

	uname, err := username.ParseUsername(*nameFl + "*" + *namespaceFl)
	if err != nil {
		return fmt.Errorf("given data produce an invalid username: %s", err)
	}

	token, err := fetchUsernameToken(*tmAddrFl, uname)
	if err != nil {
		return fmt.Errorf("cannot fetch token: %s", err)
	}

	raw, err := json.MarshalIndent(token, "", "\t")
	if err != nil {
		return fmt.Errorf("cannot json serialize token information: %s", err)
	}
	_, err = output.Write(raw)
	return err
}

func fetchUsernameToken(serverURL string, uname username.Username) (*username.Token, error) {
	resp, err := http.Get(serverURL + "/abci_query?path=%22/usernames%22&data=%22" + uname.String() + "%22")
	if err != nil {
		return nil, fmt.Errorf("cannot fetch: %s", err)
	}
	defer resp.Body.Close()

	var payload struct {
		Result struct {
			Response struct {
				Value []byte
			}
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("cannot decode response: %s", err)
	}
	var values weaveapp.ResultSet
	if err := values.Unmarshal(payload.Result.Response.Value); err != nil {
		return nil, fmt.Errorf("cannot unmarshal values: %s", err)
	}

	switch n := len(values.Results); {
	case n == 0:
		return nil, errors.New("username not found")
	case n == 1:
		// All good.
	default:
		return nil, fmt.Errorf("expected single result, got %d", n)
	}

	var token username.Token
	if err := token.Unmarshal(values.Results[0]); err != nil {
		return nil, fmt.Errorf("cannot unmarshal token: %s", err)
	}
	return &token, nil
}