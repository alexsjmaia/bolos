package main

import (
	"bolos/mensagens"
	"bolos/rotinas"
	"fmt"
)

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
