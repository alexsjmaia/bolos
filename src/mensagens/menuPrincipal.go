package mensagens

import (
	"fmt"
	"time"
)

// MenuPrincipal é o menu principal do programa, ee apenas exibe as opções disponiveis em sistema
func MenuPrincipal() {
	data_atual := time.Now().Format("02/01/2006")
	hora_atual := time.Now().Format("15:04:05")

	fmt.Print(data_atual, " - ", hora_atual, "\n")

	fmt.Println("1 - Começa Venda")
	fmt.Println("2 - Cadastra Produto")
	fmt.Println("3 - Altera Produto/Preço")
	fmt.Println("4 - Relatório de produtos cadastrados")
	fmt.Println("5 - Relatório de vendas")
	fmt.Println("0 - Sai do sistema")
}
