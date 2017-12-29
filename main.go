package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/cappyzawa/go-tiny-crud/router"
)

func main() {
	r := router.NewRouter()
	r.Logger.Fatal(r.Start(":18081"))
}
