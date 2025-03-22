package mensagens

import "fmt"

// ProdutoNaoEncontrado esta tela é exibida quando o sistema não acha o produto no banco de dados,
// esta chamada vem da tela de vendas
func ProdutoNaoEncontrado(codigo int) {
	fmt.Println("*******************************")
	fmt.Printf("*** Código %d não cadastrado ***\n", codigo)
	fmt.Println("*******************************")
}
