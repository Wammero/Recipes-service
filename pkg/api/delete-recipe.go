package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (api *api) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, exists := vars["id"]
	if !exists {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	err := api.db.DeleteRecipe(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Ok"))
}
