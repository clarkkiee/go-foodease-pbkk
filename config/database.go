package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	var  (
		host = os.Getenv("DB_HOST")
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbName = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
			PreferSimpleProtocol: true,
		},
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}