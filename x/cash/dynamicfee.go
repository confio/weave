/*

DynamicFeeDecorator is an enhanced version the basic FeeDecorator with better
handling of transaction errors and ability to deduct/enforce app-specific fees.

The business logic is:
1. If a transaction fee < min fee, or a transaction fee cannot be paid, reject
   it with an error.
2. Run the transaction.
3. If a transaction processing results in an error, revert all transaction
   changes and charge only the min fee.

TODO: If a transaction succeeded, but requested a RequiredFee higher than paid
fee, revert all transaction changes and refund all but the min fee, returning
an error.

If a transaction succeeded, and at least RequiredFee was paid, everything is
committed and we return success

It also embeds a checkpoint inside, so in the typical application stack:

	cash.NewFeeDecorator(authFn, ctrl),
	utils.NewSavepoint().OnDeliver(),

can be replaced by

	cash.NewDynamicFeeDecorator(authFn, ctrl),

As with FeeDecorator, all deducted fees are send to the collector, whose
address is configured via gconf package.

*/

package cash

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/gconf"
	"github.com/iov-one/weave/x"
)

type DynamicFeeDecorator struct {
	auth x.Authenticator
	ctrl Controller
}

var _ weave.Decorator = DynamicFeeDecorator{}

// NewDynamicFeeDecorator returns a DynamicFeeDecorator with the given
// minimum fee, and all collected fees going to a default address.
func NewDynamicFeeDecorator(auth x.Authenticator, ctrl Controller) DynamicFeeDecorator {
	return DynamicFeeDecorator{
		auth: auth,
		ctrl: ctrl,
	}
}

// Check verifies and deducts fees before calling down the stack
func (d DynamicFeeDecorator) Check(ctx weave.Context, store weave.KVStore, tx weave.Tx, next weave.Checker) (res weave.CheckResult, err error) {
	fee, payer, cache, err := d.prepare(ctx, store, tx)
	if err != nil {
		return res, errors.Wrap(err, "cannot prepare")
	}

	defer func() {
		if err == nil {
			cache.Write()
			res.GasPayment += toPayment(fee)
		} else {
			cache.Discard()
			_ = d.chargeMinimalFee(store, payer)
		}
	}()

	if err := d.chargeFee(cache, payer, fee); err != nil {
		return res, errors.Wrap(err, "cannot charge fee")
	}
	return next.Check(ctx, cache, tx)
}

// Deliver verifies and deducts fees before calling down the stack
func (d DynamicFeeDecorator) Deliver(ctx weave.Context, store weave.KVStore, tx weave.Tx, next weave.Deliverer) (res weave.DeliverResult, err error) {
	fee, payer, cache, err := d.prepare(ctx, store, tx)
	if err != nil {
		return res, errors.Wrap(err, "cannot prepare")
	}

	defer func() {
		if err == nil {
			cache.Write()
		} else {
			cache.Discard()
			_ = d.chargeMinimalFee(store, payer)
		}
	}()

	if err := d.chargeFee(cache, payer, fee); err != nil {
		return weave.DeliverResult{}, errors.Wrap(err, "cannot charge fee")
	}
	return next.Deliver(ctx, cache, tx)
}

func (d DynamicFeeDecorator) chargeFee(store weave.KVStore, src weave.Address, amount x.Coin) error {
	if amount.IsZero() {
		return nil
	}
	dest := gconf.Address(store, GconfCollectorAddress)
	return d.ctrl.MoveCoins(store, src, dest, amount)
}

// chargeMinimalFee deduct an anty span fee from a given account.
func (d DynamicFeeDecorator) chargeMinimalFee(store weave.KVStore, src weave.Address) error {
	fee := gconf.Coin(store, GconfMinimalFee)

	if fee.IsZero() {
		return nil
	}

	// If a fee has no currency set any ticker is accepted.
	if fee.Ticker == "" {
		switch funds, err := d.ctrl.Balance(store, src); {
		case err == nil:
			// Any ticker will do.
			fee.Ticker = funds[0].Ticker
		case errors.Is(errors.ErrNotFound, err):
			return errors.ErrInsufficientAmount.New("no funds")
		default:
			return errors.Wrap(err, "cannot introspect account balance")
		}
	}

	if err := d.chargeFee(store, src, fee); err != nil {
		return errors.Wrap(err, "cannot charge mimal fee")
	}
	return nil
}

// prepare is all shared setup between Check and Deliver. It computes the fee
// for the transaction, ensures that the payer is authenticated and prepares
// the database transaction.
func (d DynamicFeeDecorator) prepare(ctx weave.Context, store weave.KVStore, tx weave.Tx) (fee x.Coin, payer weave.Address, cache weave.KVCacheWrap, err error) {
	finfo, err := d.extractFee(ctx, tx, store)
	if err != nil {
		return fee, payer, cache, errors.Wrap(err, "cannot extract fee")
	}
	// Dererefence the fees (handling nil).
	if pfee := finfo.GetFees(); pfee != nil {
		fee = *pfee
	}
	payer = finfo.GetPayer()

	// Verify we have access to the money.
	if !d.auth.HasAddress(ctx, payer) {
		err := errors.ErrUnauthorized.New("fee payer signature missing")
		return fee, payer, cache, err
	}

	// Ensure we can execute subtransactions (see check on utils.Savepoint).
	cstore, ok := store.(weave.CacheableKVStore)
	if !ok {
		err = errors.ErrInternal.New("need cachable kvstore")
		return fee, payer, cache, err
	}
	cache = cstore.CacheWrap()
	return fee, payer, cache, nil
}

// this returns the fee info to deduct and the error if incorrectly set
func (d DynamicFeeDecorator) extractFee(ctx weave.Context, tx weave.Tx, store weave.KVStore) (*FeeInfo, error) {
	var finfo *FeeInfo
	ftx, ok := tx.(FeeTx)
	if ok {
		payer := x.MainSigner(ctx, d.auth).Address()
		finfo = ftx.GetFees().DefaultPayer(payer)
	}

	txFee := finfo.GetFees()
	if x.IsEmpty(txFee) {
		minFee := gconf.Coin(store, GconfMinimalFee)
		if minFee.IsZero() {
			return finfo, nil
		}
		return nil, errors.ErrInsufficientAmount.New("zero transaction fee is not allowed")
	}

	if err := finfo.Validate(); err != nil {
		return nil, errors.Wrap(err, "invalid fee")
	}

	minFee := gconf.Coin(store, GconfMinimalFee)
	// If the minimum fee has no currency accept anything.
	if minFee.Ticker == "" {
		minFee.Ticker = txFee.Ticker
	}
	if !txFee.SameType(minFee) {
		return nil, x.ErrInvalidCurrency.Newf("min fee is %s and tx fee is %s", minFee.Ticker, txFee.Ticker)

	}
	if !txFee.IsGTE(minFee) {
		return nil, errors.ErrInsufficientAmount.Newf("transaction fee less than minimum: %v", txFee)
	}
	return finfo, nil
}
