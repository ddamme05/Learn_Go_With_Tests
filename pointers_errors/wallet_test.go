package pointers_errors

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()

	if result != expected {
		t.Errorf("got %s want %s", result, expected)
	}
}

func assertNoError(t testing.TB, result error) {
	t.Helper()
	if result != nil {
		t.Errorf("got an error didn't want one")
	}
}

func assertError(t testing.TB, result, expected error) {
	t.Helper()
	if result == nil {
		t.Error("wanted an error but did not get one!")
	}

	if !errors.Is(result, expected) {
		t.Errorf("got %q, want %q", result, expected)
	}
}
