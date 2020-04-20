package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/iov-one/weave/x/escrow"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	bnsd "github.com/iov-one/weave/cmd/bnsd/app"
	"github.com/iov-one/weave/cmd/bnsd/x/username"
	"github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/x/cash"
	"github.com/iov-one/weave/x/multisig"
)

type Out struct {
	Username []tokenFormat         `json:"username"`
	Wallets  []cash.GenesisAccount `json:"cash"`
	Escrow   []escrowFormat        `json:"escrow"`
	Contract []contractFormat      `json:"contract"`
}

type tokenFormat struct {
	Username string
	Targets  []username.BlockchainAddress
	Owner    weave.Address
}

type escrowFormat struct {
	Source      weave.Address  `json:"source"`
	Arbiter     weave.Address  `json:"arbiter"`
	Destination weave.Address  `json:"destination"`
	Timeout     weave.UnixTime `json:"timeout"`
	Amount      []*coin.Coin   `json:"amount"`
	Address     weave.Address  `json:"address"`
}

type contractFormat struct {
	Participants        []*multisig.Participant `json:"participants"`
	ActivationThreshold multisig.Weight         `json:"activation_threshold"`
	AdminThreshold      multisig.Weight         `json:"admin_threshold"`
	Address             weave.Address           `json:"address"`
}

func cmdGenerateJson(input io.Reader, output io.Writer, args []string) error {
	fl := flag.NewFlagSet("", flag.ExitOnError)
	fl.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), `
Export state data. Pipe-in app version as input.`)
		fl.PrintDefaults()
	}
	var (
		dbFl = fl.String("db", env("BNSD_DB_DIR", os.ExpandEnv("$HOME")+"/.bns"),
			"bnsd database directory")
		outFl = fl.String("out", "./dump.json",
			"dump output directory")
	)

	fl.Parse(args)

	// If the given reader is providing a stat information (ie os.Stdin)
	// then check if the data is being piped. That should prevent us from
	// waiting for a data on a reader that no one ever writes to.
	if s, ok := input.(stater); ok {
		if info, err := s.Stat(); err == nil {
			isPipe := (info.Mode() & os.ModeCharDevice) == 0
			if !isPipe {
				return io.EOF
			}
		}
	}

	i, err := ioutil.ReadAll(input)
	if err != nil {
		return fmt.Errorf("cannot version input: %s", err)

	}
	version, err := strconv.Atoi(string(i))
	if err != nil {
		return fmt.Errorf("cannot convert input to string: %s", err)
	}

	// validate db
	dbPath := filepath.Join(*dbFl, "bns.db")
	_, err = os.Stat(dbPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("db file does not exists: %s", err)
	}

	// create db store
	kv, err := bnsd.CommitKVStore(dbPath)
	if err != nil {
		return fmt.Errorf("cannot initialize bnsd commit store: %s", err)
	}
	store := app.NewCommitStore(kv)
	// set db version/height
	err = kv.LoadVersion(int64(version))
	if err != nil {
		return fmt.Errorf("cannot load db version: %s", err)
	}

	// create output file
	outFile, err := os.Create(*outFl)
	if err != nil {
		return fmt.Errorf("cannot create output file: %s", err)
	}

	usernames, err := extractUsername(store)
	if err != nil {
		return fmt.Errorf("cannot extract usernames: %s", err)
	}
	escrows, err := extractEscrow(store)
	if err != nil {
		return fmt.Errorf("cannot extract escrows: %s", err)
	}
	contracts, err := extractContracts(store)
	if err != nil {
		return fmt.Errorf("cannot extract contracts: %s", err)
	}
	wallets, err := extractWallets(store)
	if err != nil {
		return fmt.Errorf("cannot extract wallets: %s", err)
	}

	outJson := Out{
		Username: usernames,
		Escrow:   escrows,
		Contract: contracts,
		Wallets:  wallets,
	}
	err = json.NewEncoder(outFile).Encode(outJson)
	if err != nil {
		return fmt.Errorf("cannot write to file: %s", err)
	}

	return nil
}

type stater interface {
	Stat() (os.FileInfo, error)
}

func extractUsername(store *app.CommitStore) ([]tokenFormat, error) {
	it := orm.IterAll("tokens")
	var out []tokenFormat
	for {
		var token username.Token
		switch key, err := it.Next(store.CheckStore(), &token); {
		case err == nil:
			out = append(out, tokenFormat{
				Username: string(key),
				Targets:  token.Targets,
				Owner:    token.Owner,
			})
		case errors.ErrIteratorDone.Is(err):
			return out, nil
		default:
			return nil, err
		}
	}
	return out, nil
}

func extractEscrow(store *app.CommitStore) ([]escrowFormat, error) {
	it := orm.IterAll("esc")
	wb := cash.NewBucket()

	var out []escrowFormat
	for {
		var e escrow.Escrow
		switch key, err := it.Next(store.CheckStore(), &e); {
		case err == nil:
			c, err := wb.Get(store.CheckStore(), key)
			if err != nil {
				return nil, err
			}
			coins := cash.AsCoins(c)
			out = append(out, escrowFormat{
				Address:     e.Address,
				Source:      e.Source,
				Arbiter:     e.Arbiter,
				Destination: e.Destination,
				Timeout:     e.Timeout,
				Amount:      coins,
			})
		case errors.ErrIteratorDone.Is(err):
			return out, nil
		default:
			return nil, err
		}
	}
	return out, nil
}

func extractContracts(store *app.CommitStore) ([]contractFormat, error) {
	it := orm.IterAll("contracts")
	var out []contractFormat
	for {
		var e multisig.Contract
		switch key, err := it.Next(store.CheckStore(), &e); {
		case err == nil:
			if err != nil {
				return nil, err
			}
			out = append(out, contractFormat{
				Participants:        e.Participants,
				ActivationThreshold: e.ActivationThreshold,
				AdminThreshold:      e.AdminThreshold,
				Address:             key,
			})
		case errors.ErrIteratorDone.Is(err):
			return out, nil
		default:
			return nil, err
		}
	}
	return out, nil
}

func extractWallets(store *app.CommitStore) ([]cash.GenesisAccount, error) {
	it := orm.IterAll("cash")
	var out []cash.GenesisAccount
	for {
		var w cash.Set
		switch key, err := it.Next(store.CheckStore(), &w); {
		case err == nil:
			s := cash.Set{
				Coins: w.Coins,
			}
			out = append(out, cash.GenesisAccount{
				Address: key,
				Set:     s,
			})
		case errors.ErrIteratorDone.Is(err):
			return out, nil
		default:
			return nil, err
		}
	}
	return out, nil
}
