package cash

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iov-one/weave"
	coin "github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInitState(t *testing.T) {
	// test data
	addr := []byte("12345678901234567890")
	coins := Set{Coins: mustCombineCoins(coin.NewCoin(100, 5, "ATM"), coin.NewCoin(50, 0, "ETH"))}
	accts := []GenesisAccount{{Address: addr, Set: coins}}

	bz, err := json.Marshal(accts)
	require.NoError(t, err)

	// hardcode
	bz2 := []byte(`[{"address":"0102030405060708090021222324252627282930",
                "coins":[{"whole":50,
                "fractional":1234567,
                "ticker":"FOO"
              }]}]`)
	coins2 := Set{Coins: mustCombineCoins(coin.NewCoin(50, 1234567, "FOO"))}
	addr2 := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30}

	// use a valid configuration so it doesn't all fail
	config := Configuration{
		CollectorAddress: weave.NewAddress([]byte("foo")),
		MinimalFee:       coin.NewCoin(0, 20, "IOV"),
	}
	rawConfig, err := json.Marshal(config)
	require.NoError(t, err)

	badConfig := Configuration{
		MinimalFee: coin.NewCoin(0, 20, "food"),
	}
	rawInvalid, err := json.Marshal(badConfig)
	require.NoError(t, err)

	cases := [...]struct {
		opts    weave.Options
		isError bool
		acct    []byte
		wallet  Set
	}{
		// no prob if no data
		0: {weave.Options{"cashconf": rawConfig}, false, nil, Set{}},
		1: {weave.Options{"foo": []byte(`"bar"`), "cashconf": rawConfig}, false, nil, Set{}},
		// unknown key
		2: {weave.Options{"foo": []byte(`[{"address": "1234"}]`), "cashconf": rawConfig}, false, nil, Set{}},
		// bad address
		3: {weave.Options{"cash": []byte(`[{"coins": 123}]`), "cashconf": rawConfig}, true, nil, Set{}},
		// get a real account
		4: {weave.Options{"cash": bz, "cashconf": rawConfig}, false, addr, coins},
		5: {weave.Options{"cash": bz2, "cashconf": rawConfig}, false, addr2, coins2},
		// enforces valid config
		6: {weave.Options{"cashconf": rawInvalid}, true, nil, Set{}},
	}

	init := Initializer{}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			kv := store.MemStore()
			bucket := NewBucket()
			err := init.FromGenesis(tc.opts, kv)
			if tc.isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if tc.acct != nil {
				acct, err := bucket.Get(kv, tc.acct)
				require.NoError(t, err)
				if assert.NotNil(t, acct) {
					assert.EqualValues(t, tc.wallet.Coins, AsCoins(acct))
				}
			}
		})
	}
}

// mustCombineCoins has one return value for tests...
func mustCombineCoins(cs ...coin.Coin) coin.Coins {
	s, err := coin.CombineCoins(cs...)
	if err != nil {
		panic(err)
	}
	return s
}
