package f_test

import (
	"testing"

	"github.com/pellared/fluentassert/f"
)

func TestOrdered(t *testing.T) {
	t.Run("Lesser", func(t *testing.T) {
		t.Run("Lesser", func(t *testing.T) {
			msg := f.Ordered(0).Lesser(1)
			assertPassed(t, msg)
		})
		t.Run("Equal", func(t *testing.T) {
			msg := f.Ordered(0).Lesser(0)
			assertFailed(t, msg, "the object is not lesser")
		})
		t.Run("Greater", func(t *testing.T) {
			msg := f.Ordered(0).Lesser(-1)
			assertFailed(t, msg, "the object is not lesser")
		})
	})

	t.Run("LesserOrEqual", func(t *testing.T) {
		t.Run("Lesser", func(t *testing.T) {
			msg := f.Ordered(0).LesserOrEqual(1)
			assertPassed(t, msg)
		})
		t.Run("Equal", func(t *testing.T) {
			msg := f.Ordered(0).LesserOrEqual(0)
			assertPassed(t, msg)
		})
		t.Run("Greater", func(t *testing.T) {
			msg := f.Ordered(0).LesserOrEqual(-1)
			assertFailed(t, msg, "the object is not lesser or equal")
		})
	})

	t.Run("GreaterOrEqual", func(t *testing.T) {
		t.Run("Lesser", func(t *testing.T) {
			msg := f.Ordered(0).GreaterOrEqual(1)
			assertFailed(t, msg, "the object is not greater or equal")
		})
		t.Run("Equal", func(t *testing.T) {
			msg := f.Ordered(0).GreaterOrEqual(0)
			assertPassed(t, msg)
		})
		t.Run("Greater", func(t *testing.T) {
			msg := f.Ordered(0).GreaterOrEqual(-1)
			assertPassed(t, msg)
		})
	})

	t.Run("Greater", func(t *testing.T) {
		t.Run("Lesser", func(t *testing.T) {
			msg := f.Ordered(0).Greater(1)
			assertFailed(t, msg, "the object is not greater")
		})
		t.Run("Equal", func(t *testing.T) {
			msg := f.Ordered(0).Greater(0)
			assertFailed(t, msg, "the object is not greater")
		})
		t.Run("Greater", func(t *testing.T) {
			msg := f.Ordered(0).Greater(-1)
			assertPassed(t, msg)
		})
	})

	t.Run("has assertions from Comparable and Obj", func(t *testing.T) {
		want := 123
		got := f.Ordered(want).FluentComparable.FluentObj.Got // type embedding done properly
		assertEqual(t, got, want)
	})
}
