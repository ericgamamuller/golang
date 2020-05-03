package main

import (
	"fmt"
	"git/course-02/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valorDoSaque float64) string
}

func main() {
	contaDoDenis := contas.ContaCorrente{}
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	fmt.Println(contaDoDenis.ObterSaldo())
}
