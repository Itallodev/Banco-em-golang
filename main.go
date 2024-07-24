package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDaClarisse := contas.ContaPoupanca{}
	contaDaClarisse.Depositar(200)
	PagarBoleto(&contaDaClarisse, 120)

	fmt.Println(contaDaClarisse.ObterSaldo())

	contaDoItallo := contas.ContaCorrente{}
	contaDoItallo.Depositar(785)
	PagarBoleto(&contaDoItallo, 288)

	fmt.Println(contaDoItallo.ObterSaldo())
}
