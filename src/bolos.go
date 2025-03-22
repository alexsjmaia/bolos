package main

import (
	"bolos/mensagens"
	"bolos/rotinas"
	"fmt"
)

// main função main responsável pelo start do programa
func main() {
	mensagens.Saudacao()

	for {
		mensagens.MenuPrincipal()

		var rotina int
		fmt.Printf("Escolha a Rotina :")
		fmt.Scan(&rotina)

		rotinas.RotinaPrincipal(rotina)
	}
}
