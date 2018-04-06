package cash

import (
	"fmt"
	"testing"

	"github.com/confio/weave"
	"github.com/confio/weave/errors"
	"github.com/confio/weave/orm"
	"github.com/confio/weave/store"
	"github.com/confio/weave/x"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type feeTx struct {
	info *FeeInfo
}

var _ weave.Tx = (*feeTx)(nil)
var _ FeeTx = feeTx{}

func (feeTx) GetMsg() (weave.Msg, error) {
	return nil, nil
}

func (f feeTx) GetFees() *FeeInfo {
	return f.info
}

func (f feeTx) Marshal() ([]byte, error) {
	return nil, errors.ErrInternal("TODO: not implemented")
}

func (f *feeTx) Unmarshal([]byte) error {
	return errors.ErrInternal("TODO: not implemented")
}

type okHandler struct{}

var _ weave.Handler = okHandler{}

func (okHandler) Check(weave.Context, weave.KVStore,
	weave.Tx) (weave.CheckResult, error) {
	return weave.CheckResult{}, nil
}

func (okHandler) Deliver(weave.Context, weave.KVStore,
	weave.Tx) (weave.DeliverResult, error) {
	return weave.DeliverResult{}, nil
}

func must(obj orm.Object, err error) orm.Object {
	if err != nil {
		panic(err)
	}
	return obj
}

func TestFees(t *testing.T) {
	var helpers x.TestHelpers

	cash := x.NewCoin(50, 0, "FOO")
	min := x.NewCoin(0, 1234, "FOO")
	addr := weave.NewAddress([]byte{1, 2, 3})
	addr2 := weave.NewAddress([]byte{3, 4, 5})
	addr3 := weave.NewAddress([]byte{0xAB})

	cases := []struct {
		signers   []weave.Address
		initState []orm.Object
		fee       *FeeInfo
		min       x.Coin
		expect    checkErr
	}{
		// no fee given, nothing expected
		0: {nil, nil, nil, x.Coin{}, noErr},
		// no fee given, something expected
		1: {nil, nil, nil, min, IsInsufficientFeesErr},
		// no signer given
		2: {nil, nil, &FeeInfo{Fees: &min}, min, errors.IsUnrecognizedAddressErr},
		// use default signer, but not enough money
		3: {
			[]weave.Address{addr},
			nil,
			&FeeInfo{Fees: &min},
			min,
			IsEmptyAccountErr,
		},
		// signer can cover min, but not pledge
		4: {
			[]weave.Address{addr},
			[]orm.Object{must(WalletWith(addr, &min))},
			&FeeInfo{Fees: &cash},
			min,
			IsInsufficientFundsErr,
		},
		// all proper
		5: {
			[]weave.Address{addr},
			[]orm.Object{must(WalletWith(addr, &cash))},
			&FeeInfo{Fees: &min},
			min,
			noErr,
		},
		// trying to pay from wrong account
		6: {
			[]weave.Address{addr},
			[]orm.Object{must(WalletWith(addr2, &cash))},
			&FeeInfo{Payer: addr2, Fees: &min},
			min,
			errors.IsUnauthorizedErr,
		},
		// can pay in any fee
		7: {
			[]weave.Address{addr},
			[]orm.Object{must(WalletWith(addr, &cash))},
			&FeeInfo{Fees: &min},
			x.NewCoin(0, 1000, ""),
			noErr,
		},
		// wrong currency checked
		8: {
			[]weave.Address{addr},
			[]orm.Object{must(WalletWith(addr, &cash))},
			&FeeInfo{Fees: &min},
			x.NewCoin(0, 1000, "NOT"),
			x.IsInvalidCurrencyErr,
		},
		// has the cash, but didn't offer enough fees
		9: {
			[]weave.Address{addr},
			[]orm.Object{must(WalletWith(addr, &cash))},
			&FeeInfo{Fees: &min},
			x.NewCoin(0, 45000, "FOO"),
			IsInsufficientFeesErr,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			auth := helpers.Authenticate(tc.signers...)
			controller := NewController(NewBucket())
			base := &Config{
				MinFee:    &tc.min,
				Collector: addr3,
			}
			h := NewFeeDecorator(auth, controller, base)

			kv := store.MemStore()
			bucket := NewBucket()
			for _, wallet := range tc.initState {
				err := bucket.Save(kv, wallet)
				require.NoError(t, err)
			}

			tx := &feeTx{tc.fee}

			_, err := h.Check(nil, kv, tx, okHandler{})
			assert.True(t, tc.expect(err), "%+v", err)
			_, err = h.Deliver(nil, kv, tx, okHandler{})
			assert.True(t, tc.expect(err), "%+v", err)
		})
	}
}
