/*
Package dbconfig reads database settings from a database.yml file (following the rails database.yml convention)
and generates a connection string for the github.com/lib/pq and github.com/go-sql-driver/mysql drivers.
*/
package dbconfig

import (
	"os"
	"strings"
)

/*
Settings returns the database settings from the database.yml file from a given application
and the corresponding application environment. The location to the database.yml file and
the environment is configured in the settings json configuration file.
If environment is NOT configured, you can set the environment variable APPLICATION_ENV (on os level).
If this is also not defined "development" is the default.
*/
func Settings(path string) map[string]string {

	jsonConf := LoadJSONConfig(path)
	dbConfig := LoadYamlConfig(jsonConf.Database_file)
	environment := jsonConf.Environment

	return dbConfig[environment]
}

/*
PostgresConnectionString returns the connection string to open a sql session
used by the github.com/lib/pq package like for example:
"host=dbserver.org password=password user=dbuser dbname=blog_production sslmode=disable"
The first parameter is the path to the database settings configuration (json) file
and the second paramater defines the sslmode.
*/
func PostgresConnectionString(sslmode string) string {
	// settings := Settings(path)

	// connection := []string{
	// 	"host=", settings["host"], " ",
	// 	"password=", settings["password"], " ",
	// 	"user=", settings["username"], " ",
	// 	"dbname=", settings["database"], " ",
	// 	"sslmode=", sslmode}

	host := os.Getenv("APP_HOST")
	password := os.Getenv("PGDB_PASSWORD")
	user := os.Getenv("PGDB_USER")
	dbname := os.Getenv("PG_DB")

	connection := []string{
		"host=", host, " ",
		"password=", password, " ",
		"user=", user, " ",
		"dbname=", dbname, " ",
		"sslmode=", sslmode}

	return strings.Join(connection, "")
}
