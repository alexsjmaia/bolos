package dataBase

import (
	"log"
)

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

	log.Fatal("Se chegou aqui não deu nada ", Produto.Id, Produto.CodigoDoProduto, Produto.DescricaoDoProduto, Produto.PrecoDeCusto, Produto.PrecoDeVenda)
}
