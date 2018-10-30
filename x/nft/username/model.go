package username

import (
	"github.com/iov-one/weave"
	"github.com/iov-one/weave/orm"
)

const (
	BucketName            = "usrnft"
	ChainAddressIndexName = "chainaddr"
	chainAddressSeparator = "*"
)

// enforce that Contract fulfils desired interface compile-time
var _ orm.CloneableData = (*UsernameToken)(nil)

// Validate enforces sigs and threshold boundaries
func (c *UsernameToken) Validate() error {
	return nil
}

// Copy makes a new Profile with the same data
func (c *UsernameToken) Copy() orm.CloneableData {
	return &UsernameToken{}
}

// ApprovalBucket is a type-safe wrapper around orm.Bucket
type UsernameTokenBucket struct {
	orm.Bucket
}

// NewApprovalBucket initializes a ApprovalBucket with default name
//
// inherit Get and Save from orm.Bucket
// add run-time check on Save
func NewUsernameTokenBucket() UsernameTokenBucket {
	bucket := orm.NewBucket(BucketName,
		orm.NewSimpleObj(nil, new(UsernameToken)))
	return UsernameTokenBucket{
		Bucket: bucket,
	}
}

// Save enforces the proper type
func (b UsernameTokenBucket) Save(db weave.KVStore, obj orm.Object) error {
	if _, ok := obj.Value().(*Approval); !ok {
		return orm.ErrInvalidObject(obj.Value())
	}
	return b.Bucket.Save(db, obj)
}
