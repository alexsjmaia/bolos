package rotinas

import (
	"bolos/alteraProduto"
	cadastroDeProdutos "bolos/cadastraProduto"
	"bolos/relatorios"
	"bolos/venda"
	"fmt"
	"os"
)

// RotinaPrincipal é a rotina responsável por chamar cada função atrelada a ela.
// Falta implementar um limpa tela para melhor visualização
func RotinaPrincipal(rotina int) {
	switch rotina {
	case 0:
		fmt.Println("Saindo do sistema...")
		os.Exit(0)
	case 1:
		fmt.Println("Começa Venda")
		venda.VendaProduto()
	case 2:
		cadastroDeProdutos.CadastraProduto()
	case 3:
		alteraProduto.AlteraProduto()
	case 4:
		relatorios.RelProdCadastrados()
	case 5:
		relatorios.RelVendas()
	default:
		fmt.Println("*** Rotina inváida! ***")
	}
}
