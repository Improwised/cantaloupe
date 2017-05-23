package services

import (
	"github.com/gocraft/dbr"
	"github.com/improwised/cantaloupe/dbconfig"
)

// CheckErr : to panic error.
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Configuration : database config.
type Configuration struct {
	DbName   string
	UserName string
}

// DbConneection : database connecting string.
func DbConneection() *dbr.Connection {
	connectionString := dbconfig.PostgresConnectionString("disable")
	conn, _ := dbr.Open("postgres", connectionString, nil)
	return conn
}

var conn *dbr.Connection

// SetupDB : setup eith database.
func SetupDB() *dbr.Session {
	if conn == nil {
		conn = DbConneection()
	}

	// create a session for each business unit of execution
	sess := conn.NewSession(nil)
	return sess
}
