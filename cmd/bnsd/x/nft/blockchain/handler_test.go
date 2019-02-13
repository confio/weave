package blockchain_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/iov-one/weave/cmd/bnsd/app"
	"github.com/iov-one/weave/store/iavl"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/cmd/bnsd/x/nft/blockchain"
	"github.com/iov-one/weave/cmd/bnsd/x/nft/ticker"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/nft"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleIssueTokenMsg(t *testing.T) {
	var helpers x.TestHelpers
	_, alice := helpers.MakeKey()
	_, bob := helpers.MakeKey()

	nft.RegisterAction(nft.DefaultActions...)

	db := store.MemStore()
	bucket := blockchain.NewBucket()
	o, _ := bucket.Create(db, bob.Address(), []byte("any_network"), nil, blockchain.Chain{MainTickerID: []byte("IOV")}, blockchain.IOV{Codec: "asd"})
	bucket.Save(db, o)
	tickerBucket := ticker.NewBucket()
	tick, _ := tickerBucket.Create(db, alice.Address(), []byte("IOV"), nil, []byte("any_network"))
	tickerBucket.Save(db, tick)

	handler := blockchain.NewIssueHandler(helpers.Authenticate(alice), nil, bucket, tickerBucket.Bucket)

	// when
	specs := []struct {
		owner, id       []byte
		details         blockchain.TokenDetails
		approvals       []nft.ActionApprovals
		expCheckError   bool
		expDeliverError bool
	}{
		{ // happy path
			owner:   alice.Address(),
			id:      []byte("other_network"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "bns"}},
		},
		{ // happy path for tendermint chain (validate autogen chainId, codec valid)
			owner:   alice.Address(),
			id:      []byte("test-chain-CnckvA"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "cosmos"}},
		},
		{ // happy path for lisk chain (validate nethash, codec valid)
			owner:   alice.Address(),
			id:      []byte("9a9813156bf1d2355da31a171e37f97dfa7568187c3fd7f9c728de8f180c19c7"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "lisk"}},
		},
		{ // valid approvals
			owner:   alice.Address(),
			id:      []byte("other_network1"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "test"}},
			approvals: []nft.ActionApprovals{{
				Action:    nft.UpdateDetails,
				Approvals: []nft.Approval{{Options: nft.ApprovalOptions{Count: nft.UnlimitedCount}, Address: bob.Address()}},
			}},
		},
		{ // invalid ticker
			owner:   alice.Address(),
			id:      []byte("other_network2"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("1OV")}, Iov: blockchain.IOV{Codec: "test", CodecConfig: `{"da": 1}`}},
			approvals: []nft.ActionApprovals{{
				Action:    nft.UpdateDetails,
				Approvals: []nft.Approval{{Options: nft.ApprovalOptions{Count: nft.UnlimitedCount}, Address: bob.Address()}},
			}},
			expDeliverError: true,
		},
		{ // unegistered ticker
			owner:           alice.Address(),
			id:              []byte("other_network3"),
			details:         blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("LSK")}, Iov: blockchain.IOV{Codec: "test", CodecConfig: `{"da": 1}`}},
			expDeliverError: true,
		},
		{ // invalid codec
			owner:   alice.Address(),
			id:      []byte("other_network4"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "1"}},
			approvals: []nft.ActionApprovals{{
				Action:    nft.UpdateDetails,
				Approvals: []nft.Approval{{Options: nft.ApprovalOptions{Count: nft.UnlimitedCount}, Address: bob.Address()}},
			}},
			expCheckError: true,
		},
		{ // invalid codec json
			owner:   alice.Address(),
			id:      []byte("other_network5"),
			details: blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "bbb", CodecConfig: "{ssdas"}},
			approvals: []nft.ActionApprovals{{
				Action:    nft.UpdateDetails,
				Approvals: []nft.Approval{{Options: nft.ApprovalOptions{Count: nft.UnlimitedCount}, Address: bob.Address()}},
			}},
			expCheckError: true,
		},
		{ // invalid approvals
			owner:           alice.Address(),
			id:              []byte("other_network6"),
			details:         blockchain.TokenDetails{Chain: blockchain.Chain{MainTickerID: []byte("IOV")}, Iov: blockchain.IOV{Codec: "test"}},
			expCheckError:   true,
			expDeliverError: true,
			approvals: []nft.ActionApprovals{{
				Action:    "12",
				Approvals: []nft.Approval{{Options: nft.ApprovalOptions{}, Address: nil}},
			}},
		},
		// todo: add other test cases when details are specified
	}

	for i, spec := range specs {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			tx := helpers.MockTx(&blockchain.IssueTokenMsg{
				Owner:     spec.owner,
				ID:        spec.id,
				Details:   spec.details,
				Approvals: spec.approvals,
			})

			// when
			cache := db.CacheWrap()
			_, err := handler.Check(nil, cache, tx)
			cache.Discard()
			if spec.expCheckError {
				require.Error(t, err)
				return
			}
			// then
			require.NoError(t, err)

			// and when delivered
			res, err := handler.Deliver(nil, db, tx)

			// then
			if spec.expDeliverError {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, res)
			assert.Equal(t, uint32(0), res.ToABCI().Code)

			// and persisted
			o, err := bucket.Get(db, spec.id)
			require.NoError(t, err)
			require.NotNil(t, o)
			u, _ := blockchain.AsBlockchain(o)
			assert.Equal(t, spec.details.Chain, u.GetChain())
			assert.Equal(t, spec.details.Iov, u.GetIov())
			// todo: verify approvals
		})
	}
}

func TestQueryTokenByName(t *testing.T) {
	var helpers x.TestHelpers
	_, alice := helpers.MakeKey()
	_, bob := helpers.MakeKey()

	nft.RegisterAction(nft.DefaultActions...)

	db := store.MemStore()
	bucket := blockchain.NewBucket()
	o1, _ := bucket.Create(db, alice.Address(), []byte("alicenet"), nil, blockchain.Chain{MainTickerID: []byte("IOV")}, blockchain.IOV{Codec: "asd"})
	bucket.Save(db, o1)
	o2, _ := bucket.Create(db, bob.Address(), []byte("bobnet"), nil, blockchain.Chain{MainTickerID: []byte("IOV")}, blockchain.IOV{Codec: "asd"})
	bucket.Save(db, o2)

	qr := weave.NewQueryRouter()
	blockchain.RegisterQuery(qr)
	// when
	h := qr.Handler("/nft/blockchains")
	require.NotNil(t, h)
	mods, err := h.Query(db, "", []byte("alicenet"))
	// then
	require.NoError(t, err)
	require.Len(t, mods, 1)

	assert.Equal(t, bucket.DBKey([]byte("alicenet")), mods[0].Key)
	got, err := bucket.Parse(nil, mods[0].Value)
	require.NoError(t, err)
	x, err := blockchain.AsBlockchain(got)
	require.NoError(t, err)
	_ = x // todo verify stored details
}

func TestDeterministicStorage(t *testing.T) {
	var helpers x.TestHelpers
	for i := 0; i < 10000; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			db := iavl.MockCommitStore()
			bucket := blockchain.NewBucket()
			tickerBucket := ticker.NewBucket()
			var owner weave.Address
			owner.UnmarshalJSON([]byte("4orppuuU/Ii3PrfL1rh7+T65vvA="))
			handler := blockchain.NewIssueHandler(helpers.AlwaysTrueAuthenticator(), nil, bucket, tickerBucket.Bucket)

			// when
			rawJson := `
 	{
    "owner": "4orppuuU/Ii3PrfL1rh7+T65vvA=",
    "id": "YWxpY2VDaGFpbjMzODM=",
    "details": {
     "chain": {},
     "iov": {
      "codec": "test",
      "codec_config": "{ \"any\" : [ \"json\", \"content\" ] }"
     }
    },
    "approvals": null
   }`

			var issueMsg blockchain.IssueTokenMsg
			require.NoError(t, json.Unmarshal([]byte(rawJson), &issueMsg))

			tx := &app.Tx{Sum: &app.Tx_IssueBlockchainNftMsg{&issueMsg}}

			_, err := handler.Deliver(nil, db.Adapter(), tx)
			require.NoError(t, err)
			hash := db.Commit().Hash
			require.Equal(t, h("792704c884c5d9163a5cb466555b927470e3b80a103de9f56d237b229d3af0cf"), hash)
		})
	}
}

func h(s string) []byte {
	v, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return v
}
