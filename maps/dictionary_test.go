package maps

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		result, _ := dictionary.Search("test")
		expected := "this is a test"

		assertStrings(t, expected, result)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, expected := dictionary.Search("unknown")

		if expected == nil {
			t.Fatal("expected an error")
		}

		assertError(t, expected, ErrNotFound)
	})
}

func assertStrings(t testing.TB, expected, result string) {
	t.Helper()

	if expected != result {
		t.Errorf("got %s want %s", result, expected)
	}
}

func assertError(t testing.TB, result, expected error) {
	t.Helper()

	if !errors.Is(expected, result) {
		t.Errorf("got error %q want %q", result, expected)
	}
}
