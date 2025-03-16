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

		descricao_produto, preco_venda := dataBase.BuscaProdutoNoBanco(codigoDigitado)

		if descricao_produto != "Erro" {
			fmt.Print("Qtd :")
			fmt.Scan(&qtdDigitada)

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

			fmt.Printf("%-4s %-40s %-10s %-10s %-14s\n", "COD", "DESCRIÇÃO", "QTD", "PREÇO", " VENDA")
			for _, produto := range produtosVendidos {
				fmt.Printf("%-4d %-40s %-10d R$ %-8.2f R$ %-8.2f\n",
					produto.CodigoDoProduto, produto.DescricaoDoProduto, produto.Qtd, produto.PrecoDeVenda, produto.TotProd)
			}

			fmt.Printf("Total Parcial R$ %.2f ", totalVenda)

		} else {
			fmt.Println("Produto não encontrado")
		}
	}
}
