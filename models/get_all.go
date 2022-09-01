package models

import "github.com/matheusgb/crud-library/database"

func GetAll() (books []Book, err error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return
	}
	defer connection.Close()

	rows, err := connection.Query(`SELECT * FROM books`)
	if err != nil {
		return
	}

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Category, &book.Author, &book.Synopsis)
		if err != nil {
			continue
		}

		books = append(books, book)
	}

	return
}
