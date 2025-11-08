package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		expected := "Hello, Chris"

		assertCorrectMessage(t, Hello("Chris"), expected)
	})

	t.Run("Saying Hello to an empty string", func(t *testing.T) {
		expected := "Hello, World"

		assertCorrectMessage(t, Hello(""), expected)
	})
}

func assertCorrectMessage(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}
