package api

import (
	"encoding/json"
	"net/http"
	"recipe/models"
)

func (api *api) AddRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	err := json.NewDecoder(r.Body).Decode(&recipe)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = api.db.AddRecipe(&recipe)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
