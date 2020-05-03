package contas

import (
	"git/course-02/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	if valorDoSaque <= 0 {
		return "Impossível sacar este valor"
	}

	if valorDoSaque <= c.saldo {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito <= 0 {
		return "Impossível depositar este valor", c.saldo
	}

	c.saldo += valorDoDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransf float64, contaDestino *ContaCorrente) (bool, string) {
	if valorDaTransf <= 0 {
		return false, "Impossível depositar este valor"
	}

	if valorDaTransf > c.saldo {
		return false, "Saldo insuficiente"
	}

	c.saldo -= valorDaTransf
	contaDestino.Depositar(valorDaTransf)
	return true, "Transferência realizada"
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
