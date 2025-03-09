package rotinas

import (
	cadastroDeProdutos "bolos/cadastraProduto"
	"fmt"
	"os"
)

func RotinaPrincipal(rotina int) {
	switch rotina {
	case 0:
		fmt.Println("Saindo do sistema...")
		os.Exit(0)
	case 1:
		fmt.Println("Começa Venda")
	case 2:
		cadastroDeProdutos.CadastraProduto()
	case 3:
		fmt.Println("Altera Produto/Preço")
	case 4:
		fmt.Println("Relatório de produtos cadastrados")
	case 5:
		fmt.Println("Relatório de Vendas")
	default:
		fmt.Println("*** Rotina inváida! ***")
	}
}
