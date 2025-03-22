package dataBase

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// RelatorioDeVendas busca as vendas no banco de dados e retorna as vendas da data escolhida pelo usuário
func RelatorioDeVendas(dataDoRelatorio string) {
	var Produto struct {
		id                int
		codigo_produto    int
		descricao_produto string
		preco_venda       float64
		qtd_venda         int
		data_venda        time.Time
		hora_venda        string
	}

	// Converter data para formato MySQL caso necessário
	parsedDate, err := time.Parse("02/01/2006", dataDoRelatorio) // Supondo que receba DD/MM/YYYY
	if err != nil {
		log.Fatal("Erro ao converter a data:", err)
	}
	dataDoRelatorio = parsedDate.Format("2006-01-02") // Formato correto para MySQL

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}
	defer db.Close()

	query := "SELECT id, codigo_produto, descricao_produto, preco_venda, qtd_venda, data_venda, hora_venda FROM venda_produto WHERE data_venda = ?"
	linhas, err := db.Query(query, dataDoRelatorio)
	if err != nil {
		log.Fatal("Erro ao criar a query", err)
	}
	defer linhas.Close()

	// Cabeçalho do relatório
	fmt.Printf("%-4s %-40s %-12s %-4s %-12s %s\n", "ID", "DESCRIÇÃO", "PREÇO", "QTD", "DATA", "HORA")

	for linhas.Next() {
		if err := linhas.Scan(&Produto.id,
			&Produto.codigo_produto,
			&Produto.descricao_produto,
			&Produto.preco_venda,
			&Produto.qtd_venda,
			&Produto.data_venda,
			&Produto.hora_venda); err != nil {
			log.Fatal("Erro ao buscar os dados no banco", err)
		}

		//descricaoLimpa serve para limpar a descrição e não dar ulo de linha ao exibir
		descricaoLimpa := strings.TrimSpace(Produto.descricao_produto)

		// Exibe os produtos vendidos de acordo com o Cabeçalho do relatório
		fmt.Printf("%-4d %-40s R$ %-9.2f %-4d %s %s\n",
			Produto.id,
			descricaoLimpa,
			Produto.preco_venda,
			Produto.qtd_venda,
			Produto.data_venda.Format("2006-01-02"),
			Produto.hora_venda,
		)
	}
}
