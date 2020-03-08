package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mysql"
)

// UserDB - defines used for Users DB type as sql.
var (
	Client *sql.DB
)

func init() {
	// Load the Env Variables.
	godotenv.Load()

	dataSourceName := fmt.Sprintf("%s:%s@/%s",
		os.Getenv("mysql_users_username"),
		os.Getenv("mysql_users_password"),
		os.Getenv("mysql_users_schema"),
	)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database successfully configured.")
}
