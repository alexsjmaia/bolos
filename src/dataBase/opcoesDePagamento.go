package dataBase

import (
	"fmt"
	"log"
)

func OpcoesDePagamento() (int, string, bool) {
	var opcaoEscolhida int

	var OpcoesPgto struct {
		Id        int
		Codigo    int
		Descricao string
	}
	var opcoes []int

	db, err := ConexaoBanco()
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados")
	}

	linha, err := db.Query("select * from opcoes_de_pagamento")
	if err != nil {
		log.Fatal("Erro ao criar o statement", err)
	}
	defer linha.Close()

	for linha.Next() {
		if err := linha.Scan(&OpcoesPgto.Id, &OpcoesPgto.Codigo, &OpcoesPgto.Descricao); err != nil {
			log.Fatal("Erro ao buscar as formas de pagamento no banco", err)
		}

		opcoes = append(opcoes, OpcoesPgto.Codigo)

		fmt.Println(OpcoesPgto.Codigo, " - ", OpcoesPgto.Descricao)
	}

	for {
		fmt.Print("Escolha uma opção :")
		fmt.Scan(&opcaoEscolhida)

		for _, codigo := range opcoes {
			if codigo == opcaoEscolhida {
				return OpcoesPgto.Codigo, OpcoesPgto.Descricao, true
			} else {
				return opcaoEscolhida, "Erro", true
			}
		}
	}
}
