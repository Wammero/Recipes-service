package api

import (
	"encoding/json"
	"net/http"
)

func (api *api) GetIngredients(w http.ResponseWriter, r *http.Request) {
	data, err := api.db.GetIngredients()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
