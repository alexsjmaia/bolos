package relatorios

import (
	"bolos/dataBase"
	"fmt"
	"time"
)

// RelProdCadastrados Exibe um cabeçalho e chama o banco de dados para exibir todos os produtos cadastrados
func RelProdCadastrados() {
	fmt.Println("Relatório de Produtos cadastrados")
	fmt.Printf("%-4s %-40s %-10s %-14s\n", "COD", "DESCRIÇÃO", "CUSTO", "VENDA")
	fmt.Println("-------------------------------------------------------------------")

	dataBase.RelProdCadastrados()

	fmt.Println("-------------------------------------------------------------------")

	time.Sleep(time.Second * 20)
}
