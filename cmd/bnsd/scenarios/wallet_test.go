package scenarios

import (
	"testing"

	"github.com/iov-one/weave/cmd/bnsd/client"
	"github.com/iov-one/weave/x"
	"github.com/stretchr/testify/require"
)

func TestSendTokens(t *testing.T) {
	emilia := client.GenPrivateKey()
	heights := make([]int64, 4)
	aNonce := client.NewNonce(bnsClient, alice.PublicKey().Address())

	walletResp, err := bnsClient.GetWallet(alice.PublicKey().Address())
	require.NotEmpty(t, walletResp.Wallet.Coins)
	for i, coin := range walletResp.Wallet.Coins {
		// send a coin from Alice to Emilia
		coin := x.Coin{
			Ticker:     coin.Ticker,
			Fractional: 0,
			Whole:      1,
		}

		seq, err := aNonce.Next()
		require.NoError(t, err)
		tx := client.BuildSendTx(alice.PublicKey().Address(), emilia.PublicKey().Address(), coin, "text tx")
		require.NoError(t, client.SignTx(tx, alice, chainID, seq))
		resp := bnsClient.BroadcastTx(tx)
		require.NoError(t, resp.IsError())
		heights[i] = resp.Response.Height
		delayForRateLimits()
	}
	walletResp, err = bnsClient.GetWallet(emilia.PublicKey().Address())
	require.NoError(t, err)
	t.Log("message", "done", "height", heights, "coins", walletResp.Wallet.Coins)
}
