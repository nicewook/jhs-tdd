package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestHelloV2(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("say 'Hello, world` when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestHelloV3(t *testing.T) {
	assertCorrectMsg := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("saying hello to Spanish people", func(t *testing.T) {
		got := HelloInternational("Chris", "spanish")
		want := "Hola, Chris"
		assertCorrectMsg(t, got, want)
	})
	t.Run("saying hello to Engilsh people", func(t *testing.T) {
		got := HelloInternational("Chris", "english")
		want := "Hello, Chris"
		assertCorrectMsg(t, got, want)
	})
	t.Run("say 'Hello, world` when an empty string is supplied", func(t *testing.T) {
		got := HelloInternational("", "")
		want := "Hello, world"
		assertCorrectMsg(t, got, want)

	})
}
