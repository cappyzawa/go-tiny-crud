package db

import (
	"database/sql"
	"os"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectMysql() *sql.DB {
	mysqlDb, err := sql.Open("mysql", os.Getenv("MYSQL_USER_NAME") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(127.0.0.1:3306)/tiny_crud")
	if err != nil {
		log.Fatal(err)
	}
	return mysqlDb
}