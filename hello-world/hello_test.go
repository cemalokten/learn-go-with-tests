package main

import "testing"

const name = "Cemal"

func TestHello(t *testing.T) {
	t.Run("Say hello to someone", func(t *testing.T) {
		got := Hello(name, "en")
		want := "Hello, Cemal"
		AssertCorrectMessage(t, want, got)
	})
	t.Run("Say hello world if empty string is provided", func(t *testing.T) {
		got := Hello("", "en")
		want := "Hello, world"
		AssertCorrectMessage(t, want, got)
	})
	t.Run("Say Hello world in Spanish", func(t *testing.T) {
		got := Hello("", "es")
		want := "Hola, world"
		AssertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello world in French", func(t *testing.T) {
		got := Hello("", "fr")
		want := "Bonjour, world"
		AssertCorrectMessage(t, got, want)
	})
}

func AssertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("Want %q, got %q", want, got)
	}
}
