package escrow

import (
	"encoding/hex"
	"encoding/json"
	"testing"

	"github.com/iov-one/weave"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenesisKey(t *testing.T) {
	const genesis = `
{
  "escrow": [
    {
      "arbiter": "foo/bar/636f6e646974696f6e64617461",
      "recipient": "C30A2424104F542576EF01FECA2FF558F5EAA61A",
      "sender": "0000000000000000000000000000000000000000",
      "timeout": "2034-11-10T23:00:00Z"
    }
  ]}`

	var opts weave.Options
	require.NoError(t, json.Unmarshal([]byte(genesis), &opts))

	db := store.MemStore()

	// when
	ini := Initializer{}
	require.NoError(t, ini.FromGenesis(opts, db))

	// then
	bucket := NewBucket()
	obj, err := bucket.Get(db, weavetest.SequenceID(1))
	require.NoError(t, err)
	require.NotNil(t, obj)
	e, ok := obj.Value().(*Escrow)
	require.True(t, ok)

	assert.Equal(t, "c30a2424104f542576ef01feca2ff558f5eaa61a", hex.EncodeToString(e.Recipient))
	assert.Equal(t, "0000000000000000000000000000000000000000", hex.EncodeToString(e.Sender))

	expArbiter := weave.NewCondition("foo", "bar", []byte("conditiondata"))
	assert.Equal(t, expArbiter, weave.Condition(e.Arbiter))
}
