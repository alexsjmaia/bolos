package dataBase

import (
	"fmt"
	"log"
)

// SalvaProduto salva oproduto digitado no banco de dados
func SalvaProduto(codigoDoProduto int, descricaoDoProduto string, precoDeCusto float64, precoDeVenda float64) {

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco", err)
	}
	defer db.Close()

	statement, err := db.Prepare("insert into cadastra_produto (codigo_produto, descricao_produto, preco_custo, preco_venda) values (?, ?, ?, ?)")

	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados", err)
	}

	resultado, err := statement.Exec(codigoDoProduto, descricaoDoProduto, precoDeCusto, precoDeVenda)

	if err != nil {
		log.Fatal("Erro ao gravar os dados no banco de dados", err)
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao capturar o Ultimo ID inserido", err, ultimoIdInserido)
	}
	fmt.Println("Produto cadastrado com Sucesso", ultimoIdInserido, codigoDoProduto, descricaoDoProduto, precoDeCusto, precoDeVenda)
}
