package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Say hello to someone", func(t *testing.T) {
		got := Hello("Cemal")
		want := "Hello, Cemal"
		AssertCorrectMessage(t, want, got)
	})
	t.Run("Say hello world if empty string is provided", func(t *testing.T) {
		got := Hello("")
		want := "Helo, world"
		AssertCorrectMessage(t, want, got)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("Want %q, got %q", want, got)
	}
}
