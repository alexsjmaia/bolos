package dataBase

import (
	"fmt"
	"log"
)

func AlteraProduto(codigoDoProduto int) {
	log.Fatal("Codigo do Produto ", codigoDoProduto)

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro co conectar ao banco de dados")
	}
	defer db.Close()

	linha, err := db.Query("SELECT * FROM cadastra_produto where id = 'codigoDoProduto' ")
	if err != nil {
		log.Fatal("Erro ao criar o statement cobrança", err)
	}
	defer linha.Close()

	var Produto struct {
		id                 int
		codigoDoProduto    int
		descricaoDoProduto string
		precoDeCusto       float64
		precoDeVenda       float64
	}

	if linha.Next() {
		if err := linha.Scan(&Produto.id, &Produto.codigoDoProduto, &Produto.descricaoDoProduto, &Produto.precoDeCusto, &Produto.precoDeVenda); err != nil {
			log.Fatal("Erro ao buscar os valores no banco", err)
		} else {
			fmt.Println(&Produto.descricaoDoProduto)
			fmt.Println(&Produto.descricaoDoProduto)
			log.Fatal("Pare aqui para ver no que deu...")
		}
	}
	log.Fatal("Se chegou aqui não deu nada")
}
