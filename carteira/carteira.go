package carteira

import "fmt"

type Bitcoin int

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(quantidade Bitcoin) {
	fmt.Printf("O endereço do saldo no Depositar é %v \n", &c.saldo)
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}
