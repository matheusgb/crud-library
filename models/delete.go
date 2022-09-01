package models

import "github.com/matheusgb/crud-library/database"

func Delete(id int) (int64, error) {
	connection, err := database.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer connection.Close()

	result, err := connection.Exec(`DELETE FROM books	WHERE id=$1`, id)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}
