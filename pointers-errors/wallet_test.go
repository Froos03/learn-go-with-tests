package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(20))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(30))
	})

	t.Run("Withdraw with Insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		err := wallet.Withdraw(Bitcoin(50))
		assertBalance(t, wallet, Bitcoin(30))
		assertError(t, err, ErrInsufficientFunds)
	})

}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("didn't want an error but got one")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.balance
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
