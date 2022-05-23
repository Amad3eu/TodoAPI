package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/aprendagolang/api-psql/models"
)

func Delete(w http.ResponseWriter, r *http.request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao remover registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}
	resp := map[string]any{
		"Error":   false,
		"Message": "registro removido com sucesso!",
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
