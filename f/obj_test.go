package f_test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestObj(t *testing.T) {
	type A struct {
		Str   string
		Bool  bool
		Slice []int
	}

	t.Run("DeepEq", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			msg := f.Obj(got).DeepEq(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
			msg := f.Obj(got).DeepEq(want)
			assertFailed(t, msg, "mismatch (-want +got):\n")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := f.Obj(got).DeepEq(nil)
			assertPassed(t, msg)
		})
	})

	t.Run("NotDeepEq", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "wrong", Bool: true, Slice: []int{1, 3}}
			msg := f.Obj(got).NotDeepEq(want)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			want := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			got := A{Str: "string", Bool: true, Slice: []int{1, 2, 3}}
			msg := f.Obj(got).NotDeepEq(want)
			assertFailed(t, msg, "the objects are equal")
		})
		t.Run("nil", func(t *testing.T) {
			var got *A
			msg := f.Obj(got).NotDeepEq(nil)
			assertFailed(t, msg, "the objects are equal")
		})
	})

	t.Run("Check", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			fn := func(x A) string {
				return ""
			}
			msg := f.Obj(A{}).Check(fn)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			fn := func(x A) string {
				return "failure"
			}
			msg := f.Obj(A{}).Check(fn)
			assertFailed(t, msg, "failure")
		})
	})

	t.Run("Should", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			pred := func(x A) bool {
				return true
			}
			msg := f.Obj(A{}).Should(pred)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			pred := func(x A) bool {
				return false
			}
			msg := f.Obj(A{}).Should(pred)
			assertFailed(t, msg, "object does not meet the predicate criteria")
		})
	})

	t.Run("ShouldNot", func(t *testing.T) {
		t.Run("Passed", func(t *testing.T) {
			pred := func(x A) bool {
				return false
			}
			msg := f.Obj(A{}).ShouldNot(pred)
			assertPassed(t, msg)
		})
		t.Run("Failed", func(t *testing.T) {
			pred := func(x A) bool {
				return true
			}
			msg := f.Obj(A{}).ShouldNot(pred)
			assertFailed(t, msg, "object meets the predicate criteria")
		})
	})
}
