package scenarios

import (
	"flag"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/cmd/bnsd/app"
	"github.com/iov-one/weave/cmd/bnsd/client"
	abci "github.com/tendermint/tendermint/abci/types"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/log"
	nm "github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/rpc/test"
	tm "github.com/tendermint/tendermint/types"
)

// application version, will be set during compilation time

const (
	startupDelay     = 100 * time.Millisecond // for tendermint setup
	testLocalAddress = "localhost:46657"
)

var (
	tendermintAddress = flag.String("address", testLocalAddress, "destination address of tendermint rpc")
	hexSeed           = flag.String("seed", "0a40d34c1970ae90acf3405f2d99dcaca16d0c7db379f4beafcfdf667b9d69ce350d27f5fb440509dfa79ec883a0510bc9a9614c3d44188881f0c5e402898b4bf3c9", "private key")
	delay             = flag.Duration("delay", time.Duration(0), "duration to wait between test cases for rate limits")
)

var (
	alice     *client.PrivateKey
	node      *nm.Node
	logger    = log.NewTMLogger(os.Stdout) //log.NewNopLogger()
	bnsClient *client.BnsClient
	chainID   string
)

func TestMain(m *testing.M) {
	flag.Parse()
	var err error
	alice, err = client.DecodePrivateKey(*hexSeed)
	if err != nil {
		logger.Error("Failed to decode private key", "cause", err)
		os.Exit(1)
	}

	if *tendermintAddress != testLocalAddress {
		bnsClient = client.NewClient(client.NewHTTPConnection(*tendermintAddress))
		chainID, err = bnsClient.ChainID()
		if err != nil {
			logger.Error("Failed to fetch chain id", "cause", err)
			os.Exit(1)
		}
		os.Exit(m.Run())
	}

	config := rpctest.GetConfig()
	config.Moniker = "SetInTestMain"
	chainID = config.ChainID()
	app, err := initApp(config, alice.PublicKey().Address())
	if err != nil {
		logger.Error("Failed to init app", "cause", err)
		os.Exit(1)
	}

	// run the app inside a tendermint instance
	node = rpctest.StartTendermint(app)
	bnsClient = client.NewClient(client.NewLocalConnection(node))
	time.Sleep(startupDelay) // wait for chain
	code := m.Run()

	// and shut down proper at the end
	node.Stop()
	node.Wait()
	os.Exit(code)
}

func delayForRateLimits() {
	time.Sleep(*delay)
}

func initApp(config *cfg.Config, addr weave.Address) (abci.Application, error) {
	bnsd, err := app.GenerateApp(config.RootDir, logger, false)
	if err != nil {
		return nil, err
	}

	// generate genesis file...
	_, err = initGenesis(config.GenesisFile(), addr)
	return bnsd, err
}

func initGenesis(filename string, addr weave.Address) (*tm.GenesisDoc, error) {
	// load genesis
	doc, err := tm.GenesisDocFromFile(filename)
	if err != nil {
		return nil, err
	}

	// set app state
	appState := fmt.Sprintf(`{
  "wallets": [
    {
      "name": "alice",
      "address": "%s",
      "coins": [
        {
          "whole": 123456789,
          "ticker": "IOV"
        },
        {
          "whole": 123456789,
          "ticker": "CASH"
        },
        {
          "whole": 123456789,
          "ticker": "ALX"
        },
        {
          "whole": 123456789,
          "ticker": "PAJA"
        }
      ]
    }
  ],
  "tokens": [
    {
      "ticker": "IOV",
      "name": "Main token of this chain",
      "sig_figs": 6
    },
    {
      "ticker": "CASH",
      "name": "Second token of this chain",
      "sig_figs": 6
    },
    {
      "ticker": "ALX",
      "name": "Third token of this chain",
      "sig_figs": 6
    },
    {
      "ticker": "PAJA",
      "name": "Mightiest token of this chain",
      "sig_figs": 9
    }
  ]
}
`, addr)
	doc.AppState = []byte(appState)
	// save file
	return doc, doc.SaveAs(filename)
}
