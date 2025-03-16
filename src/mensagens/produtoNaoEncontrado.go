package mensagens

import "fmt"

func ProdutoNaoEncontrado(codigo int) {
	fmt.Println("*******************************")
	fmt.Printf("*** Código %d não cadastrado ***\n", codigo)
	fmt.Println("*******************************")
}
