package dataBase

import (
	"log"
)

func BuscaProdutoNoBanco(codigoDigitado int) (string, float64) {
	var Produto struct {
		id                int
		codigo_produto    int
		descricao_produto string
		preco_custo       float64
		preco_venda       float64
	}

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	linha, err := db.Query("select * from cadastra_produto where codigo_produto = ?", codigoDigitado)
	if err != nil {
		log.Fatal("Erro ao criar o statement", err)
	}
	defer linha.Close()

	for linha.Next() {
		if err := linha.Scan(&Produto.id, &Produto.codigo_produto, &Produto.descricao_produto, &Produto.preco_custo, &Produto.preco_venda); err != nil {
			log.Fatal("Erro ao buscar a entrada do veiculo no banco", err)
		}
	}
	return Produto.descricao_produto, Produto.preco_venda
}
