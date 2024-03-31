package main

import (
	"fmt"

	"github.com/PedroVMB/API-GO-CELEBRITIES/database"
	"github.com/PedroVMB/API-GO-CELEBRITIES/models"
	"github.com/PedroVMB/API-GO-CELEBRITIES/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
