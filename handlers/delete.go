package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/matheusgb/crud-library/models"
)

func Delete(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error parse id: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(id)
	if err != nil {
		log.Printf("Delete error: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: deleted %d records", rows)
	}

	response := map[string]any{
		"Error":   false,
		"Message": "Successfully deleted!",
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
