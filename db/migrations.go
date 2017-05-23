package main

import (
	"os"
	// "fmt"
	// "strings"
	"database/sql"

	"github.com/DavidHuie/gomigrate"
	"github.com/improwised/cantaloupe/dbconfig"
	_ "github.com/lib/pq"
)

func main() {
	var connectionString string
	var db *sql.DB

	if os.Getenv("GO_ENV2") == "" {
		os.Setenv("GO_ENV2", "production")
	}

	connectionString = dbconfig.PostgresConnectionString("disable")
	db, _ = sql.Open("postgres", connectionString)

	migrator, err2 := gomigrate.NewMigrator(db, gomigrate.Postgres{}, "../db/migrations")
	if err2 != nil {
		panic(err2)
	}
	migrator.Migrate()
	// }
	// ************ uncomment following code to show connected user's info  **********
	// settings := dbconfig.Settings(connectionString)
	// connection := []string{
	//   "host=", settings["host"], " ",
	//   "password=", settings["password"], " ",
	//   "user=", settings["username"], " ",
	//   "dbname=", settings["database"], " ",
	//   }
	// fmt.Println(strings.Join(connection, ""))
	// **********************************************************************
	defer db.Close()
}
