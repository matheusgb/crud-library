package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/matheusgb/crud-library/models"
)

func Get(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		log.Printf("Error parse id: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	book, err := models.Get(id)
	if err != nil {
		log.Printf("Update error: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(book)
}
