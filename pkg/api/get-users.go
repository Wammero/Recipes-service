package api

import (
	"encoding/json"
	"net/http"
)

func (api *api) GetUsers(w http.ResponseWriter, r *http.Request) {
	data, err := api.db.GetUsers()

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
