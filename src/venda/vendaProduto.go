package venda

import (
	calculacobranca "bolos/calculaCobranca"
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
	var TotalVenda float64
	var codigoDigitado, qtdDigitada int
	var horaDaVenda string

	for {
		fmt.Print("\nCódigo :")
		fmt.Scan(&codigoDigitado)

		fmt.Println("Código Digitado : ", codigoDigitado)

		if codigoDigitado == 0 {
			opcaoDePagamento, descricao := dataBase.OpcoesDePagamento()

			if opcaoDePagamento == 1 {
				calculacobranca.CalcularCobranca(TotalVenda)
				fmt.Println("A opção escolhida foi ", opcaoDePagamento, descricao)
			}

			for _, produto := range ProdutosVendidos {
				horaDaVenda = dataBase.SalvaVenda(
					produto.CodigoDoProduto,
					produto.DescricaoDoProduto,
					produto.PrecoDeVenda,
					produto.Qtd,
					opcaoDePagamento,
					descricao,
				)
			}
			fmt.Println("Venda Gravada com sucesso", horaDaVenda)
			return
		}

		descricao_produto, preco_venda := dataBase.BuscaProdutoNoBanco(codigoDigitado)

		if descricao_produto != "Erro" {
			fmt.Print("Qtd :")
			fmt.Scan(&qtdDigitada)

			descricaoLimpa := strings.TrimSpace(descricao_produto)
			totProd := preco_venda * float64(qtdDigitada)
			TotalVenda += float64(totProd)

			produto := Produto{
				CodigoDoProduto:    codigoDigitado,
				DescricaoDoProduto: descricaoLimpa,
				PrecoDeVenda:       preco_venda,
				Qtd:                qtdDigitada,
				TotProd:            totProd,
			}

			ProdutosVendidos = append(ProdutosVendidos, produto)

			fmt.Printf("%-4s %-40s %-10s %-10s %-14s\n", "COD", "DESCRIÇÃO", "QTD", "PREÇO", " VENDA")

			listaProdutos()

		} else {
			mensagens.ProdutoNaoEncontrado(codigoDigitado)
		}
	}
}

func listaProdutos() {
	for _, produto := range ProdutosVendidos {
		fmt.Printf("%-4d %-40s %-10d R$ %-8.2f R$ %-8.2f\n",
			produto.CodigoDoProduto, produto.DescricaoDoProduto, produto.Qtd, produto.PrecoDeVenda, produto.TotProd)
		TotalVenda += (produto.PrecoDeVenda * float64(produto.Qtd))
	}
	totalParcial(TotalVenda)
}

func totalParcial(TotalVenda float64) {
	fmt.Printf("Total Parcial R$ %.2f ", TotalVenda)
}
