package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/matheusgb/crud-library/configs"
)

func OpenConnection() (*sql.DB, error) {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	stringConnection := fmt.Sprintf("user=%s password=%s sslmode=disable", config.DBUser, config.DBPassword)

	connection, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	err = connection.Ping()

	return connection, err
}
