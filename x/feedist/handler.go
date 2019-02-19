package feedist

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
	"github.com/iov-one/weave/x"
)

const (
	newRevenueCost    = 0
	distributeCost    = 0
	updateRevenueCost = 0
)

// RegisterQuery registers feedlist buckets for querying.
func RegisterQuery(qr weave.QueryRouter) {
	NewRevenueBucket().Register("revenues", qr)
}

// CashController allows to manage coins stored by the accounts without the
// need to directly access the bucket.
// Required functionality is implemented by the x/cash extension.
type CashController interface {
	Balance(weave.KVStore, weave.Address) (x.Coins, error)
	MoveCoins(weave.KVStore, weave.Address, weave.Address, x.Coin) error
}

// RegisterRoutes registers handlers for feedlist message processing.
func RegisterRoutes(r weave.Registry, auth x.Authenticator, ctrl CashController) {
	bucket := NewRevenueBucket()
	r.Handle(pathNewRevenueMsg, &newRevenueHandler{
		auth:   auth,
		bucket: bucket,
		ctrl:   ctrl,
	})
	r.Handle(pathDistributeMsg, &distributeHandler{
		auth:   auth,
		bucket: bucket,
		ctrl:   ctrl,
	})
	r.Handle(pathUpdateRevenueMsg, &updateRevenueHandler{
		auth:   auth,
		bucket: bucket,
		ctrl:   ctrl,
	})
}

type newRevenueHandler struct {
	auth   x.Authenticator
	bucket *RevenueBucket
	ctrl   CashController
}

func (h *newRevenueHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.CheckResult, error) {
	var res weave.CheckResult
	if _, err := h.validate(ctx, db, tx); err != nil {
		return res, err
	}
	res.GasAllocated += newRevenueCost
	return res, nil
}

func (h *newRevenueHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.DeliverResult, error) {
	var res weave.DeliverResult
	msg, err := h.validate(ctx, db, tx)
	if err != nil {
		return res, err
	}

	obj, err := h.bucket.Create(db, &Revenue{
		Admin:      msg.Admin,
		Recipients: msg.Recipients,
	})
	if err != nil {
		return res, err
	}
	res.Data = obj.Key()
	return res, nil
}

func (h *newRevenueHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*NewRevenueMsg, error) {
	rmsg, err := tx.GetMsg()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get message")
	}
	msg, ok := rmsg.(*NewRevenueMsg)
	if !ok {
		return nil, errors.InvalidMsgErr.New("unknown transaction type")
	}
	if err := msg.Validate(); err != nil {
		return msg, err
	}
	return msg, nil
}

type distributeHandler struct {
	auth   x.Authenticator
	bucket *RevenueBucket
	ctrl   CashController
}

func (h *distributeHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.CheckResult, error) {
	var res weave.CheckResult
	if _, err := h.validate(ctx, db, tx); err != nil {
		return res, err
	}
	res.GasAllocated += distributeCost
	return res, nil
}

func (h *distributeHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.DeliverResult, error) {
	var res weave.DeliverResult
	msg, err := h.validate(ctx, db, tx)
	if err != nil {
		return res, err
	}

	rev, err := h.bucket.GetRevenue(db, msg.RevenueID)
	if err != nil {
		return res, err
	}

	if err := distribute(db, h.ctrl, RevenueAccount(msg.RevenueID), rev.Recipients); err != nil {
		return res, errors.Wrap(err, "cannot distribute")
	}
	return res, nil
}

func (h *distributeHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*DistributeMsg, error) {
	rmsg, err := tx.GetMsg()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get message")
	}
	msg, ok := rmsg.(*DistributeMsg)
	if !ok {
		return nil, errors.InvalidMsgErr.New("unknown transaction type")
	}
	if err := msg.Validate(); err != nil {
		return msg, err
	}
	if _, err := h.bucket.GetRevenue(db, msg.RevenueID); err != nil {
		return nil, errors.Wrap(err, "cannot get revenue")
	}
	return msg, nil
}

type updateRevenueHandler struct {
	auth   x.Authenticator
	bucket *RevenueBucket
	ctrl   CashController
}

func (h *updateRevenueHandler) Check(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.CheckResult, error) {
	var res weave.CheckResult
	if _, err := h.validate(ctx, db, tx); err != nil {
		return res, err
	}
	res.GasAllocated += updateRevenueCost
	return res, nil
}

func (h *updateRevenueHandler) Deliver(ctx weave.Context, db weave.KVStore, tx weave.Tx) (weave.DeliverResult, error) {
	var res weave.DeliverResult
	msg, err := h.validate(ctx, db, tx)
	if err != nil {
		return res, err
	}

	rev, err := h.bucket.GetRevenue(db, msg.RevenueID)
	if err != nil {
		return res, err
	}

	// Before updating the revenue all funds must be distributed. Only a
	// revenue with no funds can be updated, so that recipients trust us.
	// Otherwise an admin could change who receives the money without the
	// previously selected recepients ever being paid.
	if err := distribute(db, h.ctrl, RevenueAccount(msg.RevenueID), rev.Recipients); err != nil {
		return res, errors.Wrap(err, "cannot distribute")
	}

	rev.Recipients = msg.Recipients
	obj := orm.NewSimpleObj(msg.RevenueID, rev)
	if err := h.bucket.Save(db, obj); err != nil {
		return res, errors.Wrap(err, "cannot save")
	}
	return res, nil
}

func (h *updateRevenueHandler) validate(ctx weave.Context, db weave.KVStore, tx weave.Tx) (*UpdateRevenueMsg, error) {
	rmsg, err := tx.GetMsg()
	if err != nil {
		return nil, errors.Wrap(err, "cannot get message")
	}
	msg, ok := rmsg.(*UpdateRevenueMsg)
	if !ok {
		return nil, errors.InvalidMsgErr.New("unknown transaction type")
	}
	if err := msg.Validate(); err != nil {
		return msg, err
	}
	return msg, nil
}

// distribute split the funds stored under the revenue address and distribute
// them according to recipients proportions. When successful, revenue account
// has no funds left after this call.
//
// It might be that not all funds can be distributed equally. Because of that a
// small leftover can remain on the revenue account after this operation.
func distribute(db weave.KVStore, ctrl CashController, source weave.Address, recipients []*Recipient) error {
	var chunks int64
	for _, r := range recipients {
		chunks += int64(r.Weight)
	}

	// Find the greatest common division for all weights. This is needed to
	// avoid leaving big fund leftovers on the source account when
	// distributing between many recipients. Or when there is only one
	// recipient with a high weight value.
	var weights []int32
	for _, r := range recipients {
		weights = append(weights, r.Weight)
	}
	div := findGcd(weights...)

	chunks = chunks / int64(div)

	balance, err := ctrl.Balance(db, source)
	switch {
	case err == nil:
		// All good.
	case errors.Is(errors.NotFoundErr, err):
		// Account does not exist, so there is are no funds to split.
		return nil
	default:
		return errors.Wrap(err, "cannot acquire revenue account balance")
	}

	// TODO normalize balance. There is no functionality that allows to
	// normalize x.Coins right now (14 Feb 2019).

	// For each currency, distribute the coins equally to the weight of
	// each recipient. This can leave small amount of coins on the original
	// account.
	for _, c := range balance {
		// Ignore those coins that have a negative value. This
		// functionality is supposed to be distributing value from
		// revenue account, not collect it. Otherwise this could be
		// used to charge the recipients instead of paying them.
		if !c.IsPositive() {
			continue
		}

		// Rest of the division can be ignored, because we transfer
		// funds to each recipients separately. Any leftover will be
		// left on the recipients account.
		one, _ := c.Divide(chunks)

		for _, r := range recipients {
			amount := one.Multiply(int64(r.Weight / div))
			// Chunk is too small to be distributed.
			if amount.IsZero() {
				continue
			}
			if err := ctrl.MoveCoins(db, source, r.Address, amount); err != nil {
				return errors.Wrap(err, "cannot move coins")
			}
		}
	}

	return nil
}

// findGcd returns greatest common division for any number of numbers.
func findGcd(values ...int32) int32 {
	switch len(values) {
	case 0:
		return 0
	case 1:
		return values[0]
	}

	res := values[0]
	for i := 1; i < len(values); i++ {
		res = gcd(res, values[i])
	}
	return res
}

// gcd returns greatest common division of two numbers.
func gcd(a, b int32) int32 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
