package models

import (
	"github.com/matheusgb/crud-library/database"
)

func Insert(book Book) (id int, err error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	query := `INSERT INTO books (title, category, author, synopsis) VALUES ($1, $2, $3, $4) RETURNING id`

	err = connection.QueryRow(query, book.Title, book.Category, book.Author, book.Synopsis).Scan(&id)

	return
}
