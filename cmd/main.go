package main

import (
	"log"
	"recipe/pkg/api"
	"recipe/repository"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gorilla/mux"
)

func main() {
	// Подключение к базе данных
	db, err := repository.New("postgresql://root:1234@0.0.0.0:5441/bd_1?sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Выполнение миграций
	migration1, err := migrate.New(
		"file://C:/Users/nazhmud1/Desktop/projects/Recipes-service/migrations",
		"postgresql://root:1234@0.0.0.0:5441/bd_1?sslmode=disable",
	)
	if err != nil {
		log.Fatalln("Cannot start migration:", err)
	}

	if err := migration1.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("No migrations to apply, continuing...")
		} else {
			log.Fatalln("Migration failed:", err)
		}
	}

	api := api.New(mux.NewRouter(), db)
	api.Handle()

	log.Fatal(api.ListenAndServe("localhost:8081"))
}
