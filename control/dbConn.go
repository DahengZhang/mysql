package control

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "dahengzhang:000000@/local_database")
	if err != nil {
		panic(err.Error())
	}
}
