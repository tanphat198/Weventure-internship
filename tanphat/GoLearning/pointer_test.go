package GoLearning

import "testing"

func TestWallet(t*testing.T){

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := Balance()

		if got != want {
			t.Errorf("got %s but want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("Want an error but didn't get one")
		}

		if got != want {
			t.Errorf("got '%s' but want '%s'",got ,want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(11)}

		Withdraw(Bitcoin(10))

		assertBalance(t , wallet, Bitcoin(1))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
