package venda

import (
	"bolos/dataBase"
	"fmt"
	"strings"
)

type Produto struct {
	CodigoDoProduto    int
	DescricaoDoProduto string
	PrecoDeVenda       float64
	Qtd                int
	TotProd            float64
}

func VendaProduto() {
	var produtosVendidos []Produto
	var totalVenda float64
	var codigoDigitado, qtdDigitada int

	for {
		fmt.Print("\nCódigo :")
		fmt.Scan(&codigoDigitado)

		fmt.Print("Qtd :")
		fmt.Scan(&qtdDigitada)

		fmt.Printf("%-20s %-10s %-15s\n", "Descrição", "Preço", "Total Parcial")

		descricao_produto, preco_venda := dataBase.BuscaProdutoNoBanco(codigoDigitado)
		descricaoLimpa := strings.TrimSpace(descricao_produto)
		totProd := preco_venda * float64(qtdDigitada)
		totalVenda += float64(totProd)

		produto := Produto{
			CodigoDoProduto:    codigoDigitado,
			DescricaoDoProduto: descricaoLimpa,
			PrecoDeVenda:       preco_venda,
			Qtd:                qtdDigitada,
			TotProd:            totProd,
		}
		produtosVendidos = append(produtosVendidos, produto)

		fmt.Printf("Total Parcial R$ %.2f ", totalVenda)

		fmt.Println("\n*******************\n")
		//fmt.Println(produtosVendidos)
		//fmt.Printf("%-1s R$ %-10.2f R$ %-40.2f\n", descricaoLimpa, preco_venda, totProd)
		for indice, produto := range produtosVendidos {
			fmt.Printf("%d, %+v\n", indice, produto)
		}
	}
}
