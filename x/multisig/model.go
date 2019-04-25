package multisig

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/orm"
)

const (
	// BucketName is where we store the contracts
	BucketName = "contracts"
	// SequenceName is an auto-increment ID counter for contracts
	SequenceName = "id"

	// Maximum value a weight value can be set to. This is uint8 capacity
	// but because we use protobuf for serialization, weight is represented
	// by uint32 and we must manually force the limit.
	maxWeightValue = 255
)

// Weight represents the strength of a signature.
type Weight int32

func (w Weight) Validate() error {
	if w < 1 {
		return errors.Wrap(errors.ErrInvalidState,
			"weight must be greater than 0")
	}
	if w > maxWeightValue {
		return errors.Wrapf(errors.ErrOverflow,
			"weight is %d and must not be greater than %d", w, maxWeightValue)
	}
	return nil
}

var _ orm.CloneableData = (*Contract)(nil)

func (c *Contract) Validate() error {
	switch n := len(c.Participants); {
	case n == 0:
		return errors.Wrap(errors.ErrInvalidModel, "no participants")
	case n > maxParticipantsAllowed:
		return errors.Wrap(errors.ErrInvalidModel, "too many participants")
	}
	return validateWeights(errors.ErrInvalidModel,
		c.Participants, c.ActivationThreshold, c.AdminThreshold)
}

func (c *Contract) Copy() orm.CloneableData {
	ps := make([]*Participant, 0, len(c.Participants))
	for _, p := range c.Participants {
		sig := make(weave.Address, len(p.Signature))
		copy(sig, p.Signature)
		ps = append(ps, &Participant{
			Signature: sig,
			Weight:    p.Weight,
		})
	}
	return &Contract{
		Participants:        ps,
		ActivationThreshold: c.ActivationThreshold,
		AdminThreshold:      c.AdminThreshold,
	}
}

// ContractBucket is a type-safe wrapper around orm.Bucket
type ContractBucket struct {
	orm.Bucket
	idSeq orm.Sequence
}

// NewContractBucket initializes a ContractBucket with default name
//
// inherit Get and Save from orm.Bucket
// add run-time check on Save
func NewContractBucket() ContractBucket {
	bucket := orm.NewBucket(BucketName,
		orm.NewSimpleObj(nil, new(Contract)))
	return ContractBucket{
		Bucket: bucket,
		idSeq:  bucket.Sequence(SequenceName),
	}
}

// Save enforces the proper type
func (b ContractBucket) Save(db weave.KVStore, obj orm.Object) error {
	if _, ok := obj.Value().(*Contract); !ok {
		return errors.WithType(errors.ErrInvalidModel, obj.Value())
	}
	return b.Bucket.Save(db, obj)
}

// Build assigns an ID to given contract instance and returns it as an orm
// Object. It does not persist the escrow in the store.
func (b ContractBucket) Build(db weave.KVStore, c *Contract) (orm.Object, error) {
	key, err := b.idSeq.NextVal(db)
	if err != nil {
		return nil, err
	}
	return orm.NewSimpleObj(key, c), nil
}

// GetContract returns a contract with given ID.
func (b ContractBucket) GetContract(store weave.KVStore, contractID []byte) (*Contract, error) {
	obj, err := b.Get(store, contractID)
	if err != nil {
		return nil, errors.Wrap(err, "bucket lookup")
	}

	if obj == nil || obj.Value() == nil {
		return nil, errors.Wrapf(errors.ErrNotFound, "contract id %q", contractID)
	}

	c, ok := obj.Value().(*Contract)
	if !ok {
		return nil, errors.Wrapf(errors.ErrInvalidModel, "invalid type: %T", obj.Value())
	}
	return c, nil
}
