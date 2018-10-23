package app

import (
	"github.com/iov-one/weave/commands"
	"github.com/iov-one/weave/crypto"
	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/cash"
	"github.com/iov-one/weave/x/sigs"
)

// Examples generates some example structs to dump out with testgen
func Examples() []commands.Example {
	wallet := &cash.Set{
		Coins: []*x.Coin{
			{Whole: 50000, Ticker: "ETH"},
			{Whole: 150, Fractional: 567000, Ticker: "BTC"},
		},
	}

	priv := crypto.GenPrivKeyEd25519()
	pub := priv.PublicKey()
	user := &sigs.UserData{
		Pubkey:   pub,
		Sequence: 17,
	}

	dst := crypto.GenPrivKeyEd25519().PublicKey().Address()
	amt := x.NewCoin(250, 0, "ETH")
	msg := &cash.SendMsg{
		Amount: &amt,
		Dest:   dst,
		Src:    pub.Address(),
		Memo:   "Test payment",
	}

	unsigned := Tx{
		Sum: &Tx_SendMsg{msg},
	}
	tx := unsigned
	sig, err := sigs.SignTx(priv, &tx, "test-123", 17)
	if err != nil {
		panic(err)
	}
	tx.Signatures = []*sigs.StdSignature{sig}

	return []commands.Example{
		{"wallet", wallet},
		{"priv_key", priv},
		{"pub_key", pub},
		{"user", user},
		{"send_msg", msg},
		{"unsigned_tx", &unsigned},
		{"signed_tx", &tx},
	}
}
