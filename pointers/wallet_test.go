package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))
		wallet.Withdraw(Bitcoin(8))

		got := wallet.Balance()
		want := Bitcoin(12)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

}
