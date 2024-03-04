package pointerserros

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		if wallet.Balance() != 0 {
			t.Errorf("error, wallet being initialized with non zero value")
		}

		want := Bitcoin(10.0)
		wallet.Deposit(want)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		initialBalance := Bitcoin(10.0)
		withdrawBalance := Bitcoin(7.0)
		want := initialBalance - withdrawBalance
		wallet.Deposit(initialBalance)
		err := wallet.Withdraw(withdrawBalance)

		assertNoError(t, err)
		assertBalance(t, wallet, want)

	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		err := wallet.Withdraw(1.0)
		assertBalance(t, wallet, 0.0)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func ExampleWallet() {
	w := Wallet{balance: Bitcoin(20.0)}
	w.Deposit(10.0)
	err := w.Withdraw(10.0)
	if err != nil && err == ErrInsufficientFunds {
		fmt.Println("insufficient funds!")
	}
	fmt.Println(w.Balance())
	//Output: 20.00 BTC
}
