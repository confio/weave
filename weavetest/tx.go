package weavetest

import (
	"encoding/binary"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/errors"
)

// Tx represents a weave transaction.
// Transaction represents a single message that is to be processed within this
// transaction.
type Tx struct {
	// Msg is the message that is to be processed by this transaction.
	Msg weave.Msg
	// Err if set is returned by any method call.
	Err error
}

var _ weave.Tx = (*Tx)(nil)

func (tx *Tx) GetMsg() (weave.Msg, error) {
	return tx.Msg, tx.Err
}

func (tx *Tx) Unmarshal(raw []byte) error {
	if tx.Msg == nil {
		return errors.Wrap(errors.ErrState, "nil message")
	}
	return tx.Msg.Unmarshal(raw)
}

func (tx *Tx) Marshal() ([]byte, error) {
	if tx.Msg == nil {
		return nil, errors.Wrap(errors.ErrState, "nil message")
	}
	return tx.Msg.Marshal()
}

// Msg represents a weave message.
// Message is a request processed by weave within a single transaction.
type Msg struct {
	// Path returned by the path method, consumed by the router.
	RoutePath string
	// Serialized represents the serialized form of this message.
	Serialized []byte
	// Err if set is returned by any method call.
	Err error
}

var _ weave.Msg = (*Msg)(nil)

func (m *Msg) Path() string {
	return m.RoutePath
}

func (m *Msg) Validate() error {
	return m.Err
}

func (m *Msg) Unmarshal(b []byte) error {
	m.Serialized = b
	return m.Err
}

func (m *Msg) Marshal() ([]byte, error) {
	return m.Serialized, m.Err
}

// SequenceID returns an ID encoded as if it was generated by the bucket
// sequence call. This function is helpful for testing sequential event
// processing and objects creation where the ID of the next saved entity can be
// determined.
func SequenceID(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}
