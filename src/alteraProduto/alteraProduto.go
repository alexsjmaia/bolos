package alteraProduto

import (
	"bolos/dataBase"
	"fmt"
)

// AlteraProduto pede o código do produto e deixa a responsabilidade de alterar para o dataBase
func AlteraProduto() {
	fmt.Println("***********************************************************************")
	fmt.Println("*** Alteração de Produto, Descrição Preço de Venda e preço de custo ***")
	fmt.Println("***********************************************************************")

	fmt.Print("Digite o código do Produto : ")
	var codigoProduto int
	fmt.Scan(&codigoProduto)

	dataBase.AlteraProduto(codigoProduto)
}
