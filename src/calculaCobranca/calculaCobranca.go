package calculacobranca

import (
	"fmt"
)

func CalcularCobranca(totalDaVenda float64) {
	var valorRecebido float64

	for {
		fmt.Print("Valor Recebido :")
		fmt.Scan(&valorRecebido)

		if valorRecebido < 0 {
			fmt.Println("Valor inválido")
		}

		if valorRecebido > totalDaVenda {
			fmt.Println("Troco R$ ", valorRecebido-totalDaVenda)
			break
		} else {
			fmt.Println("Falta R$ ", totalDaVenda-valorRecebido)
		}
	}
}
