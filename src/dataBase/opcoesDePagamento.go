package dataBase

import (
	"fmt"
	"log"
)

func OpcoesDePagamento() (int, string) {
	var opcaoEscolhida int
	var descricaoOpcaoEscolhida string

	var OpcoesPgto struct {
		Id        int
		Codigo    int
		Descricao string
	}
	var opcoes []string

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

		opcoes = append(opcoes, OpcoesPgto.Descricao)

		fmt.Println(OpcoesPgto.Codigo, " - ", OpcoesPgto.Descricao)
	}

	for {
		fmt.Print("Escolha uma opção :")
		fmt.Scan(&opcaoEscolhida)

		if opcaoEscolhida > 0 && opcaoEscolhida <= len(opcoes) {
			descricaoOpcaoEscolhida = opcoes[opcaoEscolhida-1]
			break
		} else {
			fmt.Println("Opção inválida!")
		}
	}
	return opcaoEscolhida, descricaoOpcaoEscolhida
}
