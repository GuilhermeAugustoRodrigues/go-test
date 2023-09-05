package database

import (
	"gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func StartDatabaseConnection() {
	connectionConfig := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionConfig))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

}

func MigrateDatabase() {
	StartDatabaseConnection()

	err = DB.AutoMigrate(&models.Student{})

	if err != nil {
		log.Panic(err)
	}
}
