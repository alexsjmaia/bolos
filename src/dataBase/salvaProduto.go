package database

import (
	"fmt"
	"log"
)

func SalvaProduto(codigoDoProduto int, descricaoDoProduto string, precoDeCusto float64, precoDeVenda float64) {

	fmt.Print("Código : ", codigoDoProduto, "\n")
	fmt.Print("Descrição : ", descricaoDoProduto)
	fmt.Print("Custo R$ ", precoDeCusto, "\n")
	fmt.Print("Preço de Venda R$ ", precoDeVenda, "\n")

	log.Fatal("Pare aqui...")
}
