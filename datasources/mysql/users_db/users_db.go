package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	//"os"
)

const (
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPassword = "mysql_users_password"
	mysqlUsersHost = "mysql_users_host"
	mysqlUsersSchema = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = "root" // os.Getenv(mysqlUsersUsername) //"root"
	password = "myagodapass" //os.Getenv(mysqlUsersPassword)
	host = "127.0.0.1:3306" //os.Getenv(mysqlUsersHost)
	schema = "users_db" //os.Getenv(mysqlUsersSchema)
)

//user, pw, host, schema_name
func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
	username,
	password,
	host,
	schema,
)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}