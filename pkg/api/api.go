package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r *mux.Router
}

func New(router *mux.Router) *api {
	return &api{r: router}
}

func (api *api) Handle() {
	api.r.HandleFunc("/add-recipe", AddRecipe).Methods(http.MethodPost)
	api.r.HandleFunc("/get-recipes", GetRecipes).Methods(http.MethodGet)
	api.r.HandleFunc("/get-users", GetUsers).Methods(http.MethodGet)
	api.r.HandleFunc("/get-ingredients", GetIngredients).Methods(http.MethodGet)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
