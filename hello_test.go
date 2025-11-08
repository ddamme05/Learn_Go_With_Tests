package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {
		expected := "Hello, Chris"
		result := Hello("Chris", "")

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Saying Hello to an empty string", func(t *testing.T) {
		expected := "Hello, World"
		result := Hello("", "")

		assertCorrectMessage(t, result, expected)
	})

	t.Run("Saying Hello in Spanish", func(t *testing.T) {
		expected := "Hola, Maria"
		result := Hello("Maria", "Spanish")

		assertCorrectMessage(t, result, expected)
	})
}

func assertCorrectMessage(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("got %q want %q", result, expected)
	}
}
