package models

import "github.com/matheusgb/crud-library/database"

func Get(id int) (book Book, err error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	row := connection.QueryRow(`SELECT * FROM books WHERE id=$1`, id)

	err = row.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis)

	return
}
