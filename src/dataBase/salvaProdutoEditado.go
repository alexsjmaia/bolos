package dataBase

import (
	"fmt"
	"log"
)

func SalvaProdutoEditado(codigo_produto int, descricao_produto string, preco_custo float64, preco_venda float64) {
	fmt.Println("*****************************************")
	fmt.Println("Dados Que chegaram")
	fmt.Println("Codigo :", codigo_produto)
	fmt.Println("Descrição :", descricao_produto)
	fmt.Println("Preço de Custo :", preco_custo)
	fmt.Println("Preço de Venda :", preco_venda)

	log.Fatal("Chamar o Banco para Salvar o Produto Editado")
}
