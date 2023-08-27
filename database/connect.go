package database

import (
	"fiber_test/configs"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDB(conf configs.Config) {
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHost,
		conf.DBPort,
		conf.DBUsername,
		conf.DBPassword,
		conf.DBDatabase,
	)

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal(err)
	}

}

func GetDBClient() *gorm.DB {
	return DB
}
