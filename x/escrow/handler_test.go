package escrow

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/app"
	"github.com/iov-one/weave/coin"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/iov-one/weave/x"
	"github.com/iov-one/weave/x/cash"
	"github.com/iov-one/weave/x/hashlock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	blockNow = time.Now().UTC()
	Timeout  = weave.AsUnixTime(blockNow.Add(2 * time.Hour))

	zeroBucket = orm.NewBucket("zero", nil)
)

// TestHandler runs a number of scenario of tx to make
// sure they work as expected.
//
// I really should get quickcheck working....
func TestHandler(t *testing.T) {

	a := weavetest.NewCondition()
	b := weavetest.NewCondition()
	c := weavetest.NewCondition()
	// d is just an observer, no role in escrow
	d := weavetest.NewCondition()

	// good
	all := mustCombineCoins(coin.NewCoin(100, 0, "FOO"),
		coin.NewCoin(10, 0, "BAR"))
	some := mustCombineCoins(coin.NewCoin(32, 0, "FOO"),
		coin.NewCoin(1, 0, "BAR"))
	remain := mustCombineCoins(coin.NewCoin(68, 0, "FOO"),
		coin.NewCoin(9, 0, "BAR"))

	escrowAddr := func(i uint64) weave.Address {
		return Condition(weavetest.SequenceID(i)).Address()
	}

	cases := map[string]struct {
		// initial balance to set
		account weave.Address
		balance []*coin.Coin
		// preparation transactions, must all succeed
		prep []action
		// tx to test
		do action
		// check if do should return an error
		isError bool
		// otherwise, a series of queries...
		queries []query
	}{
		"simplest test, sending money we have creates an escrow": {
			a.Address(),
			all,
			nil, // no prep, just one action
			createAction(a, b, c, all, ""),
			false,
			[]query{
				// verify escrow is stored
				{
					"/escrows", "", weavetest.SequenceID(1), false,
					[]orm.Object{
						NewEscrow(weavetest.SequenceID(1), a.Address(), b.Address(), c, all, Timeout, ""),
					},
					NewBucket().Bucket,
				},
				// bank deducted from sender
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						cash.NewWallet(a.Address()),
					},
					cash.NewBucket().Bucket,
				},
				// and added to escrow
				{"/wallets", "", escrowAddr(1), false,
					[]orm.Object{
						mo(cash.WalletWith(escrowAddr(1), all...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
		"partial send, default sender taken from permissions": {
			a.Address(),
			all,
			nil, // no prep, just one action
			createAction(a, b, c, some, ""),
			false,
			[]query{
				// verify escrow is stored
				{
					"/escrows", "", weavetest.SequenceID(1), false,
					[]orm.Object{
						NewEscrow(weavetest.SequenceID(1), a.Address(), b.Address(), c, some, Timeout, ""),
					},
					NewBucket().Bucket,
				},
				// make sure sender index works
				{
					"/escrows/sender", "", a.Address(), false,
					[]orm.Object{
						NewEscrow(weavetest.SequenceID(1), a.Address(), b.Address(), c, some, Timeout, ""),
					},
					NewBucket().Bucket,
				},
				// make sure recipient index works
				{
					"/escrows/recipient", "", b.Address(), false,
					[]orm.Object{
						NewEscrow(weavetest.SequenceID(1), a.Address(), b.Address(), c, some, Timeout, ""),
					},
					NewBucket().Bucket,
				},
				// make sure arbiter index works
				{
					"/escrows/arbiter", "", c, false,
					[]orm.Object{
						NewEscrow(weavetest.SequenceID(1), a.Address(), b.Address(), c, some, Timeout, ""),
					},
					NewBucket().Bucket,
				},
				// make sure wrong query misses
				{
					"/escrows/arbiter", "", b, false, nil, NewBucket().Bucket,
				},
				// others id are empty
				{
					"/escrows", "", weavetest.SequenceID(2), false, nil, zeroBucket,
				},
				// bank deducted from sender
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(a.Address(), remain...)),
					},
					cash.NewBucket().Bucket,
				},
				// and added to escrow
				{"/wallets", "", escrowAddr(1), false,
					[]orm.Object{
						mo(cash.WalletWith(escrowAddr(1), some...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
		"cannot send money we don't have": {
			a.Address(),
			some,
			nil, // no prep, just one action
			createAction(a, b, c, all, ""),
			true,
			nil,
		},
		"cannot send money from other account": {
			a.Address(),
			all,
			nil, // no prep, just one action
			action{
				// note permission is not the sender!
				perms: []weave.Condition{b},
				msg:   NewCreateMsg(a.Address(), b.Address(), c, some, Timeout, ""),
			},
			true,
			nil,
		},
		"cannot set timeout in the past": {
			a.Address(),
			all,
			nil, // no prep, just one action
			action{
				perms: []weave.Condition{a},
				// defaults to sender!
				msg:       NewCreateMsg(nil, b.Address(), c, all, weave.AsUnixTime(blockNow.Add(-2*time.Hour)), ""),
				blockTime: Timeout.Time().Add(-time.Hour),
			},
			true,
			nil,
		},
		"arbiter can successfully release all": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, all, "")},
			action{
				perms: []weave.Condition{c},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
			},
			false,
			[]query{
				// verify escrow is deleted
				{
					"/escrows", "", weavetest.SequenceID(1), false, nil, zeroBucket,
				},
				// escrow is empty
				{"/wallets", "", escrowAddr(1), false,
					[]orm.Object{
						cash.NewWallet(escrowAddr(1)),
					},
					cash.NewBucket().Bucket,
				},
				// sender is broke
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						cash.NewWallet(a.Address()),
					},
					cash.NewBucket().Bucket,
				},
				// recipient has bank
				{"/wallets", "", b.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(b.Address(), all...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
		"sender can successfully release part": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, all, "hello")},
			action{
				perms: []weave.Condition{a},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
					Amount:   some,
				},
			},
			false,
			[]query{
				// verify escrow balance is updated
				{
					"/escrows", "", weavetest.SequenceID(1), false,
					[]orm.Object{
						NewEscrow(weavetest.SequenceID(1), a.Address(), b.Address(), c, remain, Timeout, "hello"),
					},
					NewBucket().Bucket,
				},
				// escrow is reduced
				{"/wallets", "", escrowAddr(1), false,
					[]orm.Object{
						mo(cash.WalletWith(escrowAddr(1), remain...)),
					},
					cash.NewBucket().Bucket,
				},
				// sender is broke
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						cash.NewWallet(a.Address()),
					},
					cash.NewBucket().Bucket,
				},
				// recipient has some money
				{"/wallets", "", b.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(b.Address(), some...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
		"recipient cannot release": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, all, "")},
			action{
				perms: []weave.Condition{b},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
			},
			true,
			nil,
		},
		"cannot release after timeout": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, all, "")},
			action{
				perms: []weave.Condition{c},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
				blockTime: Timeout.Time().Add(time.Hour),
			},
			true,
			nil,
		},
		//"successful return after expired (can be done by anyone)": {
		//	a.Address(),
		//	all,
		//	[]action{createAction(a, b, c, all, "")},
		//	action{
		//		perms: []weave.Condition{a},
		//		msg: &ReturnEscrowMsg{
		//			EscrowId: weavetest.SequenceID(1),
		//		},
		//		height: Timeout + 1,
		//	},
		//	false,
		//	[]query{
		//		// verify escrow is deleted
		//		{
		//			"/escrows", "", weavetest.SequenceID(1), false, nil, zeroBucket,
		//		},
		//		// escrow is empty
		//		{"/wallets", "", escrowAddr(1), false,
		//			[]orm.Object{
		//				cash.NewWallet(escrowAddr(1)),
		//			},
		//			cash.NewBucket().Bucket,
		//		},
		//		// sender recover all his money
		//		{"/wallets", "", a.Address(), false,
		//			[]orm.Object{
		//				mo(cash.WalletWith(a.Address(), all...)),
		//			},
		//			cash.NewBucket().Bucket,
		//		},
		//		// recipient doesn't get paid
		//		{"/wallets", "", b.Address(), false, nil,
		//			cash.NewBucket().Bucket,
		//		},
		//	},
		//},
		//"cannot return before timeout": {
		//	a.Address(),
		//	all,
		//	[]action{createAction(a, b, c, all, "")},
		//	action{
		//		perms: []weave.Condition{a},
		//		msg: &ReturnEscrowMsg{
		//			EscrowId: weavetest.SequenceID(1),
		//		},
		//		height: Timeout - 1,
		//	},
		//	true,
		//	nil,
		//},
		"we update the arbiter and then make sure the new actors are used": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, some, ""),
				{
					perms: []weave.Condition{c},
					// c hands off to d
					msg: &UpdateEscrowPartiesMsg{
						EscrowId: weavetest.SequenceID(1),
						Arbiter:  d,
					},
				}},
			action{
				// new arbiter can resolve
				perms: []weave.Condition{d},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
			},
			false,
			[]query{
				// verify escrow is deleted (resolved)
				{
					"/escrows", "", weavetest.SequenceID(1), false, nil, zeroBucket,
				},
				// bank deducted from sender
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(a.Address(), remain...)),
					},
					cash.NewBucket().Bucket,
				},
				// and added to recipient
				{"/wallets", "", b.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(b.Address(), some...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
		"after update, original arbiter cannot resolve": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, some, ""),
				{
					perms: []weave.Condition{c},
					// c hands off to d
					msg: &UpdateEscrowPartiesMsg{
						EscrowId: weavetest.SequenceID(1),
						Arbiter:  d,
					},
				}},
			action{
				// original arbiter can no longer resolve
				perms: []weave.Condition{c},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
			},
			true,
			nil,
		},
		"cannot update without proper permissions": {
			a.Address(),
			all,
			[]action{createAction(a, b, c, some, "")},
			action{
				perms: []weave.Condition{a},
				msg: &UpdateEscrowPartiesMsg{
					EscrowId: weavetest.SequenceID(1),
					Arbiter:  a,
				},
			},
			true,
			nil,
		},
		//"cannot update parties after timeout": {
		//	a.Address(),
		//	all,
		//	[]action{createAction(a, b, c, some, "")},
		//	action{
		//		perms: []weave.Condition{a},
		//		msg: &UpdateEscrowPartiesMsg{
		//			EscrowId: weavetest.SequenceID(1),
		//			Sender:   d,
		//		},
		//		height: Timeout + 100,
		//	},
		//	true,
		//	nil,
		//},
		"cannot claim escrow twice": {
			a.Address(),
			all,
			[]action{
				createAction(a, b, c, all, ""),
				{
					perms: []weave.Condition{c},
					msg: &ReleaseEscrowMsg{
						EscrowId: weavetest.SequenceID(1),
					},
				},
			},
			action{
				perms: []weave.Condition{c},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
			},
			true,
			[]query{
				// verify escrow is deleted
				{
					"/escrows", "", weavetest.SequenceID(1), false, nil, zeroBucket,
				},
				// escrow is empty
				{"/wallets", "", escrowAddr(1), false,
					[]orm.Object{
						cash.NewWallet(escrowAddr(1)),
					},
					cash.NewBucket().Bucket,
				},
				// sender is broke
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						cash.NewWallet(a.Address()),
					},
					cash.NewBucket().Bucket,
				},
				// recipient has cash
				{"/wallets", "", b.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(b.Address(), all...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
		//"return overpaid amount and delete escrow": {
		//	a.Address(),
		//	mustCombineCoins(coin.NewCoin(2, 0, "FOO")),
		//	[]action{
		//		createAction(a, b, c, mustCombineCoins(coin.NewCoin(1, 0, "FOO")), ""),
		//		{
		//			perms: []weave.Condition{a},
		//			msg: &cash.SendMsg{
		//				Src:    a.Address(),
		//				Dest:   escrowAddr(1),
		//				Amount: &coin.Coin{Whole: 1, Ticker: "FOO"},
		//			},
		//		},
		//	},
		//	action{
		//		perms: []weave.Condition{a},
		//		msg: &ReturnEscrowMsg{
		//			EscrowId: weavetest.SequenceID(1),
		//		},
		//		height: Timeout + 1,
		//	},
		//	false,
		//	[]query{
		//		// verify escrow is deleted
		//		{
		//			"/escrows", "", weavetest.SequenceID(1), false, nil, zeroBucket,
		//		},
		//		// escrow is empty
		//		{"/wallets", "", escrowAddr(1), false,
		//			[]orm.Object{
		//				cash.NewWallet(escrowAddr(1)),
		//			},
		//			cash.NewBucket().Bucket,
		//		},
		//		// sender recover all his money
		//		{"/wallets", "", a.Address(), false,
		//			[]orm.Object{
		//				mo(cash.WalletWith(a.Address(), mustCombineCoins(coin.NewCoin(2, 0, "FOO"))...)),
		//			},
		//			cash.NewBucket().Bucket,
		//		},
		//		// recipient doesn't get paid
		//		{"/wallets", "", b.Address(), false, nil,
		//			cash.NewBucket().Bucket,
		//		},
		//	},
		//},
		"release overpaid amount and delete escrow": {
			a.Address(),
			mustCombineCoins(coin.NewCoin(2, 0, "FOO")),
			[]action{
				createAction(a, b, c, mustCombineCoins(coin.NewCoin(1, 0, "FOO")), ""),
				{
					perms: []weave.Condition{a},
					msg: &cash.SendMsg{
						Src:    a.Address(),
						Dest:   escrowAddr(1),
						Amount: &coin.Coin{Whole: 1, Ticker: "FOO"},
					},
				},
			},
			action{
				perms: []weave.Condition{c},
				msg: &ReleaseEscrowMsg{
					EscrowId: weavetest.SequenceID(1),
				},
			},
			false,
			[]query{
				// verify escrow is deleted
				{
					"/escrows", "", weavetest.SequenceID(1), false, nil, zeroBucket,
				},
				// escrow is empty
				{"/wallets", "", escrowAddr(1), false,
					[]orm.Object{
						cash.NewWallet(escrowAddr(1)),
					},
					cash.NewBucket().Bucket,
				},
				// sender is broke
				{"/wallets", "", a.Address(), false,
					[]orm.Object{
						cash.NewWallet(a.Address()),
					},
					cash.NewBucket().Bucket,
				},
				// recipient has bank
				{"/wallets", "", b.Address(), false,
					[]orm.Object{
						mo(cash.WalletWith(b.Address(), mustCombineCoins(coin.NewCoin(2, 0, "FOO"))...)),
					},
					cash.NewBucket().Bucket,
				},
			},
		},
	}

	bank := cash.NewBucket()
	ctrl := cash.NewController(bank)
	auth := authenticator()
	// create handler objects and query objects
	router := app.NewRouter()
	RegisterRoutes(router, auth, ctrl)
	cash.RegisterRoutes(router, auth, ctrl)
	qr := weave.NewQueryRouter()
	cash.RegisterQuery(qr)
	RegisterQuery(qr)

	for descr, tc := range cases {
		t.Run(descr, func(t *testing.T) {
			db := store.MemStore()

			// set initial data
			acct, err := cash.WalletWith(tc.account, tc.balance...)
			require.NoError(t, err)
			err = bank.Save(db, acct)
			require.NoError(t, err)

			// do delivertx
			for j, p := range tc.prep {
				// try check
				cache := db.CacheWrap()
				_, err = router.Check(p.ctx(), cache, p.tx())
				require.NoError(t, err, "%d", j)
				cache.Discard()

				// then perform
				_, err = router.Deliver(p.ctx(), db, p.tx())
				require.NoError(t, err, "%d", j)
			}

			_, err = router.Deliver(tc.do.ctx(), db, tc.do.tx())
			if tc.isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// run through all queries
			for k, q := range tc.queries {
				q.check(t, db, qr, "%d", k)
			}
		})
	}
}

func createAction(sender, rcpt, arbiter weave.Condition, amount coin.Coins, memo string) action {
	return action{
		perms: []weave.Condition{sender},
		msg:   NewCreateMsg(sender.Address(), rcpt.Address(), arbiter, amount, Timeout, memo),
	}
}

// MinusCoins returns a-b
func MinusCoins(a, b coin.Coins) (coin.Coins, error) {
	// TODO: add coins.Negative...
	minus := b.Clone()
	for _, m := range minus {
		m.Whole *= -1
		m.Fractional *= -1
	}
	return a.Combine(minus)
}

func MustMinusCoins(t *testing.T, a, b coin.Coins) coin.Coins {
	remain, err := MinusCoins(a, b)
	require.NoError(t, err)
	return remain
}

func MustAddCoins(t *testing.T, a, b coin.Coins) coin.Coins {
	res, err := a.Combine(b)
	require.NoError(t, err)
	return res
}

// TestAtomicSwap combines hash and escrow to perform
// atomic swap...
//
// we tested timeout above, this is just about claiming
func TestAtomicSwap(t *testing.T) {
	// a and b want to do a swap
	a := weavetest.NewCondition()
	b := weavetest.NewCondition()
	// c is just an observer, no role in escrow
	c := weavetest.NewCondition()

	foo := mustCombineCoins(coin.NewCoin(500, 0, "FOO"))
	lilFoo := mustCombineCoins(coin.NewCoin(77, 0, "FOO"))
	leftFoo := MustMinusCoins(t, foo, lilFoo)
	bar := mustCombineCoins(coin.NewCoin(1100, 0, "BAR"))
	lilBar := mustCombineCoins(coin.NewCoin(250, 0, "BAR"))
	leftBar := MustMinusCoins(t, bar, lilBar)

	cases := []struct {
		// initial values
		aInit, bInit coin.Coins
		// amount we wish to swap
		aSwap, bSwap coin.Coins
		// arbiter, same on both
		arbiter weave.Condition
		// preimage used in claim
		preimage []byte
		// does the release cause an error?
		isError        bool
		aFinal, bFinal coin.Coins
	}{
		// good preimage
		0: {
			foo, bar,
			lilFoo, lilBar,
			hashlock.PreimageCondition([]byte{7, 8, 9}),
			[]byte{7, 8, 9},
			false,
			// the coins were properly released
			MustAddCoins(t, leftFoo, lilBar),
			MustAddCoins(t, leftBar, lilFoo),
		},
		// bad preimage
		1: {
			foo, bar,
			lilFoo, lilBar,
			hashlock.PreimageCondition([]byte{1, 2, 3}),
			[]byte("foo"),
			true,
			// money stayed in escrow
			leftFoo,
			leftBar,
		},
	}

	bank := cash.NewBucket()
	ctrl := cash.NewController(bank)

	setBalance := func(t *testing.T, db weave.KVStore, addr weave.Address, coins coin.Coins) {
		acct, err := cash.WalletWith(addr, coins...)
		require.NoError(t, err)
		err = bank.Save(db, acct)
		require.NoError(t, err)
	}
	checkBalance := func(t *testing.T, db weave.KVStore, addr weave.Address) coin.Coins {
		acct, err := bank.Get(db, addr)
		require.NoError(t, err)
		coins := cash.AsCoins(acct)
		return coins
	}

	// use both context auth and hashlock auth
	auth := x.ChainAuth(authenticator(), hashlock.Authenticate{})
	setAuth := authenticator().SetConditions

	// route the escrow commands, and wrap with the hashlock
	// middleware
	r := app.NewRouter()
	RegisterRoutes(r, auth, ctrl)
	h := weavetest.Decorate(r, hashlock.NewDecorator())

	ctx := weave.WithHeight(context.Background(), 500)
	ctx = weave.WithBlockTime(ctx, blockNow)
	for i, tc := range cases {
		t.Run(fmt.Sprintf("case-%d", i), func(t *testing.T) {
			// start with the balance
			db := store.MemStore()
			setBalance(t, db, a.Address(), tc.aInit)
			setBalance(t, db, b.Address(), tc.bInit)

			// make sure this works at all....
			abal := checkBalance(t, db, a.Address())
			require.Equal(t, tc.aInit, abal)
			bbal := checkBalance(t, db, b.Address())
			require.Equal(t, tc.bInit, bbal)

			// create the offer
			one := NewCreateMsg(a.Address(), b.Address(), tc.arbiter, tc.aSwap, Timeout, "")
			aCtx := setAuth(ctx, a)
			res, err := h.Deliver(aCtx, db, &weavetest.Tx{Msg: one})
			require.NoError(t, err)
			esc1 := res.Data

			// this is the response
			two := NewCreateMsg(b.Address(), a.Address(), tc.arbiter, tc.bSwap, Timeout, "")
			bCtx := setAuth(ctx, b)
			res, err = h.Deliver(bCtx, db, &weavetest.Tx{Msg: two})
			require.NoError(t, err)
			esc2 := res.Data

			// now try to execute them, c with hashlock....
			resCtx := setAuth(ctx, c)
			resTx1 := PreimageTx{
				Tx:       &weavetest.Tx{Msg: &ReleaseEscrowMsg{EscrowId: esc1}},
				Preimage: tc.preimage,
			}
			_, err = h.Deliver(resCtx, db, resTx1)
			if tc.isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
			resTx2 := PreimageTx{
				Tx:       &weavetest.Tx{Msg: &ReleaseEscrowMsg{EscrowId: esc2}},
				Preimage: tc.preimage,
			}
			_, err = h.Deliver(resCtx, db, resTx2)
			if tc.isError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			// make sure final balance is proper....
			abal = checkBalance(t, db, a.Address())
			require.Equal(t, tc.aFinal, abal)
			bbal = checkBalance(t, db, b.Address())
			require.Equal(t, tc.bFinal, bbal)
		})
	}
}

// --- cut and paste from hashlock/decorator_test.go :(

// PreimageTx fulfills the HashKeyTx interface to satisfy the decorator
type PreimageTx struct {
	weave.Tx
	Preimage []byte
}

var _ hashlock.HashKeyTx = PreimageTx{}
var _ weave.Tx = PreimageTx{}

func (p PreimageTx) GetPreimage() []byte {
	return p.Preimage
}

//-------------------------------------------------
// specific helpers for these tests

type action struct {
	perms []weave.Condition
	msg   weave.Msg

	// if not zero, overwrites blockTime function for timeout
	blockTime time.Time
}

func (a action) tx() weave.Tx {
	return &weavetest.Tx{Msg: a.msg}
}

func (a action) ctx() weave.Context {
	ctx := context.Background()
	if !a.blockTime.IsZero() {
		ctx = weave.WithBlockTime(ctx, a.blockTime)
	} else {
		ctx = weave.WithBlockTime(ctx, blockNow)
	}
	return authenticator().SetConditions(ctx, a.perms...)
}

// authenticator returns a default for all tests...
// clean this up?
func authenticator() *weavetest.CtxAuth {
	return &weavetest.CtxAuth{Key: "auth"}
}

// how to do a query... TODO: abstract this??

type query struct {
	path     string
	mod      string
	data     []byte
	isError  bool
	expected []orm.Object
	bucket   orm.Bucket
}

func (q query) check(t *testing.T, db weave.ReadOnlyKVStore,
	qr weave.QueryRouter, msg ...interface{}) {

	h := qr.Handler(q.path)
	require.NotNil(t, h)
	mods, err := h.Query(db, q.mod, q.data)
	if q.isError {
		require.Error(t, err)
		return
	}
	require.NoError(t, err)
	if assert.Equal(t, len(q.expected), len(mods), msg...) {
		for i, ex := range q.expected {
			// make sure keys match
			key := q.bucket.DBKey(ex.Key())
			assert.Equal(t, key, mods[i].Key)

			// parse out value
			got, err := q.bucket.Parse(nil, mods[i].Value)
			require.NoError(t, err)
			assert.EqualValues(t, ex.Value(), got.Value(), msg...)
		}
	}
}

// for test, panics if cannot convert to model....
func objToModel(obj orm.Object) (weave.Model, error) {
	// ugh, we need the full on length...
	key := obj.Key()
	val := obj.Value()
	// this is soo ugly....
	if _, ok := val.(*Escrow); ok {
		key = NewBucket().DBKey(key)
	} else if _, ok := val.(*cash.Set); ok {
		key = cash.NewBucket().DBKey(key)
	}
	bz, err := val.Marshal()
	return weave.Model{Key: key, Value: bz}, err
}

// mo = must object... takes (Object, error) result and
// convert to Object or panic
func mo(obj orm.Object, err error) orm.Object {
	if err != nil {
		panic(err)
	}
	return obj
}
