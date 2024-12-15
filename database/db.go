package database

import (
	"log"

	"github.com/lmparedes1977/Curso_CI_2-Projeto_base/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
