package dataBase

import (
	"fmt"
	"log"
	"time"
)

func SalvaVenda(codigoDoProduto int, descricaoDoProduto string, precoDeCusto float64, precoDeVenda float64) {

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco", err)
	}
	defer db.Close()
	data_venda := time.Now().Format("02/01/2006")
	hora_venda := time.Now().Format("15:04:05")

	statement, err := db.Prepare("insert into cadastra_produto (codigo_produto, descricao_produto, preco_custo, preco_venda, data_venda, hora_venda) values (?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados", err)
	}

	resultado, err := statement.Exec(codigoDoProduto, descricaoDoProduto, precoDeCusto, precoDeVenda, data_venda, hora_venda)

	if err != nil {
		log.Fatal("Erro ao gravar os dados no banco de dados", err)
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao capturar o Ultimo ID inserido", err, ultimoIdInserido)
	}
	fmt.Println("Produto cadastrado com Sucesso", ultimoIdInserido, codigoDoProduto, descricaoDoProduto, precoDeCusto, precoDeVenda, data_venda, hora_venda)
}
