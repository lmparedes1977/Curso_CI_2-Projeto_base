package main

import (
	"github.com/lmparedes1977/Curso_CI_2-Projeto_base/database"
	"github.com/lmparedes1977/Curso_CI_2-Projeto_base/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
