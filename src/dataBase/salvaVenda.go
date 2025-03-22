package dataBase

import (
	"fmt"
	"log"
	"time"
)

// SalvaVenda salva a venda digitana na tela de venda de produto
func SalvaVenda(CodigoDoProduto int, DescricaoDoProduto string,
	PrecoDeVenda float64, Qtd_Venda int, opcaoDePagamento int, descricao string) string {

	fmt.Println(CodigoDoProduto, DescricaoDoProduto, PrecoDeVenda, Qtd_Venda, opcaoDePagamento, descricao)

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar com o banco", err)
	}
	defer db.Close()

	data_venda := time.Now().Format("2006-01-02") // Formato YYYY-MM-DD (Correto para MySQL)
	hora_venda := time.Now().Format("15:04:05")

	statement, err := db.Prepare("insert into venda_produto (codigo_produto, descricao_produto, preco_venda, qtd_venda, data_venda, hora_venda) values (?, ?, ?, ?, ?, ?)")

	if err != nil {
		log.Fatal("Erro ao conectar no banco de dados", err)
	}

	resultado, err := statement.Exec(CodigoDoProduto, DescricaoDoProduto, PrecoDeVenda, Qtd_Venda, data_venda, hora_venda)

	if err != nil {
		log.Fatal("Erro ao gravar os dados no banco de dados", err)
	}

	ultimoIdInserido, err := resultado.LastInsertId()
	if err != nil {
		fmt.Println("Erro ao capturar o Ultimo ID inserido", err, ultimoIdInserido)
	}
	return hora_venda
}
