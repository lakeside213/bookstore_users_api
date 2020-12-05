package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUsername = "root"
	mysqlPassword = "pass"
	mysqlHost     = "localhost:3306"
	mysqlSchema   = "users_db"
)

var (
	UsersDB *sql.DB

	uname  = os.Getenv(mysqlUsername)
	pass   = os.Getenv(mysqlPassword)
	host   = os.Getenv(mysqlHost)
	schema = os.Getenv(mysqlSchema)
)

func init() {
	// datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", "pass", "localhost:3306", "users_db")
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlUsername, mysqlPassword, mysqlHost, mysqlSchema)
	var err error
	UsersDB, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = UsersDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database configured")
}
