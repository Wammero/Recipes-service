package main

import (
	"Recipes_service/cmd/migrate"
	"Recipes_service/internal/config"
	"Recipes_service/internal/repository"
	"log"
)

func main() {
	connstr := loadDatabaseConnectionString()
	repo := connectToDatabase(connstr)
	defer repo.Close()

	applyMigrations(connstr)
}

func loadDatabaseConnectionString() string {
	connstr, err := config.LoadDatabaseConnectionString()
	if err != nil {
		log.Fatalf("Failed to load database connection string: %v", err)
	}
	return connstr
}

func connectToDatabase(connstr string) *repository.PGRepo {
	repo, err := repository.New(connstr)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database!")
	return repo
}

func applyMigrations(connstr string) {
	m, err := migrate.CallMigrations(connstr)
	if err != nil {
		log.Fatalf("Ошибка при создании мигратора: %v", err)
	}

	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Println("Миграции уже применены")
		} else {
			log.Fatalf("Ошибка при применении миграции: %v", err)
		}
	} else {
		log.Println("Миграции успешно применены")
	}
}
