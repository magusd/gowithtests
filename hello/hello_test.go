package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Lobs", "")
		want := "Hello, Lobs!"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello World! when no name is provided", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world' in Spanish", func(t *testing.T) {
		got := Hello("Lobs", "es")
		want := "Hola, Lobs!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello world' in French", func(t *testing.T) {
		got := Hello("Lobs", "fr")
		want := "Bonjour, Lobs!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
