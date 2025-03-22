package dataBase

import (
	"fmt"
	"log"
	"strings"
)

// RelProdCadastrados busca todos os produtos cadastrados no banco e exibe
func RelProdCadastrados() {

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

	linhas, err := db.Query("select * from cadastra_produto")
	if err != nil {
		log.Fatal("Erro ao criar o statement", err)
	}
	defer linhas.Close()

	for linhas.Next() {
		if err := linhas.Scan(&Produto.id, &Produto.codigo_produto, &Produto.descricao_produto, &Produto.preco_custo, &Produto.preco_venda); err != nil {
			log.Fatal("Erro ao buscar os produtos no banco", err)
		}

		// descricaoLimpa limpa a descrição para que não haja pulo de linha no terminal
		descricaoLimpa := strings.TrimSpace(Produto.descricao_produto)

		// Imprime os produtos na tela os parametros são para ficarem alinhados a tela
		fmt.Printf("%-4d %-40s R$ %-9.2f R$ %-11.2f\n", Produto.id, descricaoLimpa, Produto.preco_custo, Produto.preco_venda)
	}
}
