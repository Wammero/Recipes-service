package api

import (
	"net/http"
	"recipe/repository"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) Handle() {
	api.r.HandleFunc("/add-recipe", api.AddRecipe).Methods(http.MethodPost)
	api.r.HandleFunc("/get-recipes", api.GetRecipes).Methods(http.MethodGet)
	api.r.HandleFunc("/get-users", api.GetUsers).Methods(http.MethodGet)
	api.r.HandleFunc("/get-ingredients", api.GetIngredients).Methods(http.MethodGet)
	api.r.HandleFunc("/delete-recipe", api.GetIngredients).Methods(http.MethodDelete)
	api.r.HandleFunc("/change-favourite", api.ChangeFavoutite).Methods(http.MethodPost)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
