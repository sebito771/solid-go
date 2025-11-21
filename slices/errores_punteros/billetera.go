package errorespunteros

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Stringer interface {
	String() string
}

type Billetera struct {
	balanc Bitcoin
}

func (b *Billetera) deposito(amount Bitcoin) {
	fmt.Printf("la direccion del balance es %p \n", &b.balanc)

	b.balanc += amount
}

func (b *Billetera) balance() Bitcoin {
	return b.balanc
}

func (bit Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", bit)
}
var saldoIsuficiente = errors.New("fondos insuficientes")
func (b *Billetera) Withdraw(amount Bitcoin) error {
	if b.balanc < amount {
		return saldoIsuficiente
	}
	b.balanc -= amount
	return nil
}
