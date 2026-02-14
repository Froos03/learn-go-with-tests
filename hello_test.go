package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("sayin hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' to people when an empty string is supplied", func(t *testing.T) {
		got:= Hello("")
		want:= "Hello World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}