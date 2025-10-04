package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to someone", func(t *testing.T) {
		got := Hello("Nick", "")
		want := "Hello, Nick!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to world if no name is empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Emil", "French")
		want := "Bonjour, Emil!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
