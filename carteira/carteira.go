package carteira

import (
	"errors"
	"fmt"
)

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Carteira struct {
	saldo Bitcoin
}

type Stringer interface {
	String() string
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}
	c.saldo -= quantidade
	return nil
}
