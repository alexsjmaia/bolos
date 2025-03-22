package dataBase

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var TAMANHO_MINIMO_DESCRICAO_PRODUTO = 3

// AlteraProduto Altera os dados do produto, Servindo tambem para alterar o preço de custo e preço de venda
func AlteraProduto(codigoDoProduto int) {

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro co conectar ao banco de dados")
	}
	defer db.Close()

	linhas, err := db.Query("SELECT * FROM cadastra_produto WHERE codigo_Produto = ?", codigoDoProduto)

	if err != nil {
		log.Fatal("Erro ao criar o statement cobrança", err)
	}
	defer linhas.Close()

	var Produto struct {
		Id                 int
		CodigoDoProduto    int
		DescricaoDoProduto string
		PrecoDeCusto       float64
		PrecoDeVenda       float64
	}

	if linhas.Next() {
		if err := linhas.Scan(&Produto.Id, &Produto.CodigoDoProduto, &Produto.DescricaoDoProduto, &Produto.PrecoDeCusto, &Produto.PrecoDeVenda); err != nil {
			log.Fatal("Erro ao buscar os valores no banco", err)
		}
	}

	reader := bufio.NewReader(os.Stdin) // Cria um leitor para capturar entradas completas

	fmt.Println("Descrição Atual do Produto :", Produto.DescricaoDoProduto)

	reader.ReadString('\n') // <-- Limpa o buffer após a leitura do código
	fmt.Printf("Digite uma nova Descrição :")

	var novaDescricaoProduto string
	var novoPrecoDeCusto float64
	var novoPrecoDeVenda float64

	novaDescricaoProduto, _ = reader.ReadString('\n') // Lê a linha inteira, incluindo espaços
	if len(novaDescricaoProduto) < TAMANHO_MINIMO_DESCRICAO_PRODUTO {
		novaDescricaoProduto = Produto.DescricaoDoProduto
	}

	fmt.Println("Preço de Custo :", Produto.PrecoDeCusto)
	fmt.Printf("Novo preço de custo :")
	fmt.Scan(&novoPrecoDeCusto)

	fmt.Println("Preço de Venda :", Produto.PrecoDeVenda)
	fmt.Printf("Novo Preço de Venda :")
	fmt.Scan(&novoPrecoDeVenda)

	SalvaProdutoEditado(Produto.CodigoDoProduto, novaDescricaoProduto, novoPrecoDeCusto, novoPrecoDeVenda)
}
