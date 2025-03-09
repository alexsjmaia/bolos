package main

import (
	"bolos/mensagens"
	"bolos/rotinas"
	"fmt"
)

func main() {
<<<<<<< HEAD
=======
	apagarEstaFuncao()
>>>>>>> 2c44758 (comitando da maquina da loja)
	mensagens.Saudacao()

	for {
		mensagens.MenuPrincipal()

		var rotina int
		fmt.Printf("Escolha a Rotina :")
		fmt.Scan(&rotina)
		rotinas.RotinaPrincipal(rotina)
	}
}
<<<<<<< HEAD
=======

func apagarEstaFuncao() {
	fmt.Println("Apague esta função, eles etá apenas testando o git")
}
>>>>>>> 2c44758 (comitando da maquina da loja)
