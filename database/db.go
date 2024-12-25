package database

import (
	"log"
	"os"
	"github.com/lmparedes1977/Curso_CI_2-Projeto_base/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	
	stringDeConexao := "host=" + os.Getenv("HOST") +
					   " user=" + os.Getenv("USER") +
					   " password=" + os.Getenv("PASSWORD") +
					   " dbname=" + os.Getenv("DBNAME") +
					   " port=" + os.Getenv("DBPORT") 
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
