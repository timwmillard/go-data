package data

import "testing"

func TestZeroString(t *testing.T) {
	var want string

	got := Zero[string]()
	if got != want {
		t.Errorf("not zero string, got %q but want %q", got, want)
	}
}

func TestFromPtrString(t *testing.T) {
	var want string

	got := Zero[string]()
	if got != want {
		t.Errorf("not zero string, got %q but want %q", got, want)
	}
}
