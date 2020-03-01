package orm

import (
	"reflect"
	"testing"

	"github.com/iov-one/weave/errors"
	"github.com/iov-one/weave/store"
	"github.com/iov-one/weave/weavetest/assert"
)

func TestLimitIterator(t *testing.T) {
	db := store.MemStore()

	b := NewSerialModelBucket("cnts", &CounterWithID{},
		WithIndexSerial("counter", lexographicCountIndex, true))

	var expected []*CounterWithID
	for i := 0; i < 30; i++ {
		c := &CounterWithID{
			Count: int64(i * 1000),
		}
		expected = append(expected, c)
	}

	for _, e := range expected {
		// make sure we point to value in array, so this PrimaryKey gets set
		err := b.Save(db, e)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	iter, err := b.IndexScan(db, "counter", nil, false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	var dest CounterWithID
	// should return iterator done when received limit lesser than 1
	limit := 0
	limitedIter := LimitIterator(iter, limit)
	err = limitedIter.LoadNext(&dest)
	if !errors.ErrIteratorDone.Is(err) {
		t.Fatalf("unexpected error: %v", err)
	}

	limit = -100
	limitedIter = LimitIterator(iter, limit)
	err = limitedIter.LoadNext(&dest)
	if !errors.ErrIteratorDone.Is(err) {
		t.Fatalf("unexpected error: %v", err)
	}

	// limit 1 should work
	iter, err = b.IndexScan(db, "counter", nil, false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	limit = 1
	limitedIter = LimitIterator(iter, limit)
	err = limitedIter.LoadNext(&dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	iter, err = b.IndexScan(db, "counter", nil, false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	limit = 10
	limitedIter = LimitIterator(iter, limit)
	err = limitedIter.LoadNext(&dest)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// reset iterator counter
	iter, err = b.IndexScan(db, "counter", nil, false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	limitedIter = LimitIterator(iter, limit)
	for i := 0; i < 10; i++ {
		if err := limitedIter.LoadNext(&dest); err != nil {
			if !errors.ErrIteratorDone.Is(err) {
				t.Fatalf("unexpected error: %s", err)
			}
		}
		if !reflect.DeepEqual(expected[i], &dest) {
			t.Errorf("values do not match, expected: %+v, got: %+v", expected[i], &dest)
		}
	}
}

func TestToSlice(t *testing.T) {
	db := store.MemStore()

	b := NewSerialModelBucket("cnts", &CounterWithID{},
		WithIndexSerial("counter", func(Object) ([]byte, error) { return []byte("all"), nil }, false))

	var expected []*CounterWithID
	for i := 0; i > 30; i++ {
		c := &CounterWithID{
			Count: int64(i * 1000),
		}
		expected = append(expected, c)
	}
	for _, e := range expected {
		// make sure we point to value in array, so this PrimaryKey gets set
		err := b.Save(db, e)
		assert.Nil(t, err)
	}

	iter, err := b.IndexScan(db, "counter", []byte("all"), false)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	var dest []*CounterWithID
	err = ToSlice(iter, &dest)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	if !reflect.DeepEqual(dest, expected) {
		t.Errorf("values do not match, expected: %+v, got: %+v", expected, dest)
	}
}
