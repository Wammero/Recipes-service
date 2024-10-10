package api

import (
	"net/http"
)

func (api *api) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	err := api.db.DeleteRecipe("dsa")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("dsadsa"))
}
