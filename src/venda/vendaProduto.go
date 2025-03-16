package venda

import (
	"bolos/dataBase"
	"bolos/mensagens"
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

var ProdutosVendidos []Produto
var TotalVenda float64

func VendaProduto() {
	var produtosVendidos []Produto
	var totalVenda float64
	var codigoDigitado, qtdDigitada int

	for {
		fmt.Print("\nCódigo :")
		fmt.Scan(&codigoDigitado)

		if codigoDigitado == '*' {
			mostraProdutos()

			//opcoesDePagamento()

			// Se o Pagamento não for aprovado retorna a tela de produtos

			// Se o Pagamento for aprovado salva no banco

			//dataBase.SalvaVenda()
		}

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

			ProdutosVendidos = append(produtosVendidos, produto)

			mostraProdutos()

		} else {
			mensagens.ProdutoNaoEncontrado(codigoDigitado)
		}
	}
}

func mostraProdutos() {
	fmt.Printf("%-4s %-40s %-10s %-10s %-14s\n", "COD", "DESCRIÇÃO", "QTD", "PREÇO", " VENDA")
	for _, produto := range ProdutosVendidos {
		fmt.Printf("%-4d %-40s %-10d R$ %-8.2f R$ %-8.2f\n",
			produto.CodigoDoProduto, produto.DescricaoDoProduto, produto.Qtd, produto.PrecoDeVenda, produto.TotProd)
	}

	fmt.Printf("Total Parcial R$ %.2f ", TotalVenda)
}
