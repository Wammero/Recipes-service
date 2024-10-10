package main

import (
	"log"
	"recipe/pkg/api"

	"github.com/gorilla/mux"
)

func main() {
	api := api.New(mux.NewRouter())
	api.Handle()
	log.Fatal(api.ListenAndServe("localhost:8080"))
}
