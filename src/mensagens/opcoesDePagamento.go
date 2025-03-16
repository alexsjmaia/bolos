package mensagens

import (
	"fmt"
	"log"
)

func OpcoesDePagamento() {
	fmt.Println("Escolha uma opção de pagamento")
	fmt.Println("1 - Dinheiro")
	fmt.Println("2 Cartão de crédito")
	fmt.Println("3 - Cartão de débito")
	fmt.Println("4 - PIX")

	log.Fatal("Escolha a opção de pagamento...")
}
