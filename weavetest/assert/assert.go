package assert

import (
	"reflect"
	"testing"

	"github.com/iov-one/weave/errors"
)

// Nil fails the test if given value is not nil.
func Nil(t testing.TB, value interface{}) {
	t.Helper()
	if !isNil(value) {
		// Use %+v so that if we are printing an error that supports
		// stack traces then a full stack trace is shown.
		t.Fatalf("want a nil value, got %+v", value)
	}
}

func isNil(value interface{}) (isnil bool) {
	if value == nil {
		return true
	}

	defer func() {
		if recover() != nil {
			isnil = false
		}
	}()

	// The argument must be a chan, func, interface, map, pointer, or slice
	// value; if it is not, IsNil panics.
	isnil = reflect.ValueOf(value).IsNil()

	return isnil
}

// Equal fails the test if two values are not equal.
func Equal(t testing.TB, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("values not equal \nwant %T %v\n got %T %v", want, want, got, got)
	}
}

// Panics will run given function and recover any panic. It will fail the test
// if given function call did not panic.
func Panics(t testing.TB, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("panic expected")
		}
	}()
	fn()
}

// FieldError ensures that given error contains the exact match of a single
// field error, tested by its type (.Is method call).
// To test that no error was found for a given field name, use `nil` as the
// match value.
func FieldError(t testing.TB, err error, fieldName string, want *errors.Error) {
	t.Helper()

	errs := errors.FieldErrors(err, fieldName)

	// This is a special case when we want no errors (nil).
	if want == nil {
		switch len(errs) {
		case 0:
			// All good.
			return
		case 1:
			t.Fatalf("expected no error, got %q", errs[0])
		default:
			for i, e := range errs {
				t.Logf("\terror %d: %q", i+1, e)
			}
			t.Fatalf("expected no error, got %d", errs)
		}
	}

	switch len(errs) {
	case 0:
		t.Fatal("no error found")
	case 1:
		if !want.Is(errs[0]) {
			t.Fatalf("unexpected error found: %q", errs[0])
		}
	default:
		t.Errorf("want one error, got %d", len(errs))
		for i, e := range errs {
			t.Logf("\terror %d: %q", i+1, e)
		}
		has := false
		for _, e := range errs {
			if want.Is(e) {
				has = true
				break
			}
		}
		if !has {
			t.Fatalf("error not found")
		}

	}
}
