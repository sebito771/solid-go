package errorespunteros

import (
	"testing"
)

func TestBillet(t *testing.T) {
	assertbalance := func(t testing.TB, billetera Billetera, want Bitcoin) {
		t.Helper()
		got := billetera.balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
          
	
	

	t.Run("deposito", func(t *testing.T) {
		billetera := Billetera{}
		billetera.deposito(Bitcoin(10))
		assertbalance(t, billetera, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		billetera := Billetera{balanc: Bitcoin(20)}
		err := billetera.Withdraw(Bitcoin(10))
		assertbalance(t, billetera, Bitcoin(10))
		if err != nil {
			t.Errorf("didn't expect an error but got one: %s", err)
		}
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		billetera := Billetera{startingBalance}

		err := billetera.Withdraw(Bitcoin(100))
        assertError(t,err,saldoIsuficiente)
		assertbalance(t, billetera, startingBalance)

		if err == nil {
			t.Error("no tienes nada pobreton")
		}
	})
}
func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}