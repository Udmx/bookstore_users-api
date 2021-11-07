package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // just import
	"log"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_username"
	mysql_users_host     = "mysql_users_host"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	schema   = os.Getenv(mysql_users_schema)
)

//if import users_db package called init
func init() {
	//user - password - host - database name
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName)) //TODO : delete because for test

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err) //try to connect
	}
	if err := Client.Ping(); err != nil {
		panic(err) //try to connect
	}
	log.Println("database successfully configure")
}
