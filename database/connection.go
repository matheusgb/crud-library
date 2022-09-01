package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/matheusgb/crud-library/configs"
)

func OpenConnection() (*sql.DB, error) {
	config := configs.GetDB()

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Database)

	connection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	err = connection.Ping()

	return connection, err
}
