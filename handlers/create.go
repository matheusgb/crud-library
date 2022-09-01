package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/matheusgb/crud-library/models"
)

func Create(writer http.ResponseWriter, request *http.Request) {
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		log.Printf("Error while decoding json: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(book)

	var response map[string]any

	if err != nil {
		response = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Insert error: %v", err),
		}
	} else {
		response = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Book inserted! ID: %d", id),
		}
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
