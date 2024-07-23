package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"

		assertMessage(t, got, want)
	})
	t.Run("test 2", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertMessage(t, got, want)
	})
	t.Run("test 3", func(t *testing.T) {
		got := Hello("Chris", "Spanish")
		want := "Hola, Chris"

		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
