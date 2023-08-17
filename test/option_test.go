package test

import (
	"testing"

	opt "github.com/countersoda/go-opt"
)

func TestOption(t *testing.T) {
	some := opt.Some(42)
	if some.Unwrap() != 42 {
		t.Fatalf("Failed: Got %v, wanted %d", some.Unwrap(), 42)
	}
	if some.IsNone() == true {
		t.Fatalf("Failed: Got %v, wanted %t", some.IsNone(), false)
	}
	if some.IsSome() == false {
		t.Fatalf("Failed: Got %v, wanted %t", some.IsNone(), true)
	}
	if !some.IsSomeAnd(func(num int) bool { return num > 20 }) {
		t.Fatalf("Failed: Got %v, wanted %t", some.IsNone(), true)
	}
	if some.Expect("should not be empty") != 42 {
		t.Fatalf("Failed: Got %v, wanted %d", some, 42)
	}
	someClone := some.Clone()
	if (&someClone).Replace(2).Unwrap() != 2 && some.Unwrap() != 42 {
		t.Fatalf("Failed: Got %v, wanted %d", someClone, 2)
	}
	none := opt.None[int]()
	if none.UnwrapOr(42) != 42 {
		t.Fatalf("Failed: Got %v, wanted %d", none.UnwrapOr(42), 42)
	}
	if none.UnwrapOrDefault() != 0 {
		t.Fatalf("Failed: Got %v, wanted %d", none.UnwrapOrDefault(), 0)
	}
	if (&none).Replace(42).Unwrap() != 42 {
		t.Fatalf("Failed: Got %v, wanted %d", none.Unwrap(), 42)
	}
	errMsg := "value should not be empty"
	none.Expect(errMsg)
	defer func() {
		if r := recover(); r != nil && r != errMsg {
			t.Fatalf("Failed: Got %v", r)
		}
	}()

}
