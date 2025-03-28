package dataBase

import (
	"fmt"
	"log"
	"time"
)

// SalvaProdutoEditado salva o produto editado no banco de dados
func SalvaProdutoEditado(codigo_produto int, descricao_produto string, preco_custo float64, preco_venda float64) {
	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}
	defer db.Close()

	_, err = db.Exec("update cadastra_produto set descricao_produto = ?, preco_custo = ?, preco_venda = ? WHERE codigo_produto = ?", descricao_produto, preco_custo, preco_venda, codigo_produto)
	if err != nil {
		log.Fatal("Erro ao atualizar o produto:", err)
	}
	fmt.Print("\n\nProduto Atualizado com sucesso\n")
	fmt.Println("Código :", codigo_produto)
	fmt.Println("Descrição :", descricao_produto)
	fmt.Println("Preço de custo R$ ", preco_custo)
	fmt.Println("Preço de venda R$ ", preco_venda)

	time.Sleep(time.Second * 5)
}
