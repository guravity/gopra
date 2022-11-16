package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper() // ヘルパーであることを明示
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to go", func(t *testing.T) {
		got := Hello("Go World", "")
		want := "Hello, Go World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello when alone", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "es")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Charlotte", "fr")
		want := "Bonjour, Charlotte"
		assertCorrectMessage(t, got, want)
	})

}
