package main

import (
	"Recipes_service/cmd/migrate"
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

	m, err := migrate.CallMigrations(connstr)
	if err != nil {
		log.Fatalf("Ошибка при создании мигратора: %v\n", err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("Миграции уже применены")
		} else {
			log.Fatalf("Ошибка при применении миграции: %v\n", err)
		}
	} else {
		log.Println("Миграции успешно применены")
	}

}
