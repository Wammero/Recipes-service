package main

import (
	"Recipes_service/internal/config"
	"Recipes_service/internal/repository"
	"log"
)

func main() {
	// Загрузить строку подключения к базе данных
	connstr, err := config.LoadDatabaseConnectionString()
	if err != nil {
		log.Fatalf("Failed to load database connection string: %v", err)
	}

	repo, err := repository.New(connstr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer repo.Close()

	log.Println("Successfully connected to the database!")
}
