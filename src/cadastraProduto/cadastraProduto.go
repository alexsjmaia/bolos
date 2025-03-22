package cadastroDeProdutos

import (
	database "bolos/dataBase"
	"bufio"
	"fmt"
	"os"
)

// CadastraProduto Solicita os dados do produto e chama o database.SalvaProduto para gravar os dados
func CadastraProduto() {
	var Produto struct {
		codigoDoProduto    int
		descricaoDoProduto string
		precoDeCusto       float64
		precoDeVenda       float64
	}

	reader := bufio.NewReader(os.Stdin) // Cria um leitor para capturar entradas completas

	fmt.Println("***************************")
	fmt.Println("*** CADASTRANDO PRODUTO ***")
	fmt.Println("***************************")

	fmt.Printf("Digite o código do Produto : ")
	fmt.Scan(&Produto.codigoDoProduto)
	reader.ReadString('\n') // <-- Limpa o buffer após a leitura do código

	fmt.Printf("Descricao do Produto : ")
	Produto.descricaoDoProduto, _ = reader.ReadString('\n') // Lê a linha inteira, incluindo espaços

	fmt.Printf("Preço de Custo R$ ")
	fmt.Scan(&Produto.precoDeCusto)

	fmt.Printf("Preço de Venda R$ ")
	fmt.Scan(&Produto.precoDeVenda)

	// Chama o database.SalvaProduto para salvar os dados do produto
	database.SalvaProduto(Produto.codigoDoProduto, Produto.descricaoDoProduto, Produto.precoDeCusto, Produto.precoDeVenda)
}
