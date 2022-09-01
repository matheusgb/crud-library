package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matheusgb/crud-library/models"
)

func List(writer http.ResponseWriter, request *http.Request) {
	books, err := models.GetAll()
	if err != nil {
		log.Printf("It was not possible get all records: %v", err)
	}

	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}
