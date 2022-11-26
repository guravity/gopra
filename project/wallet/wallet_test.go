package wallet

import (
	// "fmt"
	"testing"
)



func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("%#v got %s want %s", wallet, got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil { // エラー内容がnilのときは、おかしい
			t.Fatal("wanted an error but didn't get one")
		}
		if got != want{ // エラー内容が期待と違う場合
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		// fmt.Printf("Deposit method, %v \n", &wallet.balance) // アドレス確認
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw too much", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
	})

}
