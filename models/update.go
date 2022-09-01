package models

import "github.com/matheusgb/crud-library/database"

func Update(id int, book Book) (int64, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	result, err := connection.Exec(`UPDATE books SET title=$1, category=$2, author=$3, synopsis=$4
	WHERE id=$5`, book.Title, book.Category, book.Author, book.Synopsis, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
