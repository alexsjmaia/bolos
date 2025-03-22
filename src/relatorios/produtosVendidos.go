package relatorios

import (
	"bolos/dataBase"
	"fmt"
	"time"
)

// RelVendas Solicita a Data em que o usuário quer ver o relatório de vendas,
// falta implementar aqui o periodo informado ao invés de uma só data
func RelVendas() {

	dataDoRelatorio := time.Now().Format("2006-01-02") // Formato YYYY-MM-DD (Correto para MySQL)

	fmt.Print("Digite a data do relatório :")
	fmt.Scan(&dataDoRelatorio)
	fmt.Println(dataDoRelatorio)

	dataBase.RelatorioDeVendas(dataDoRelatorio)
}
