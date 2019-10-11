package leap

import "testing"

func TestLeapYear(t *testing.T) {
	t.Run("1941 is not divisible by 4", func(t *testing.T) {
		got := IsLeapYear(1941)
		want := false

		if got != want {
			t.Fatalf("got %t want %t", got, want)
		}
	})

	t.Run("2000 is divisible by 4, divisible by 400", func(t *testing.T) {
		got := IsLeapYear(2000)
		want := true

		if got != want {
			t.Fatalf("got %t want %t", got, want)
		}
	})

	t.Run("1996 is divisible by 4, not divisible by 100", func(t *testing.T) {
		got := IsLeapYear(1996)
		want := true

		if got != want {
			t.Fatalf("got %t want %t", got, want)
		}
	})

	t.Run("2100 is divisible by 100, not divisible by 400", func(t *testing.T) {
		got := IsLeapYear(2100)
		want := false

		if got != want {
			t.Fatalf("got %t want %t", got, want)
		}
	})
}
