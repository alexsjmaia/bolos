package venda

import (
	"bolos/dataBase"
	"fmt"
	"log"
)

func VendaProduto() {
	var codigoDigitado int

	fmt.Println("Código :")
	fmt.Scan(&codigoDigitado)

	descricao, precoDevenda := dataBase.BuscaProdutoNoBanco(codigoDigitado)

	log.Fatal("Chegamos na tela de venda de produto", descricao, precoDevenda)

}
