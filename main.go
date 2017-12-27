package main

import (
	"net/http"
	"log"
	"github.com/cappyzawa/go-tiny-crud/api"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/users", api.ListUsers)
	router.GET("/users/:id", api.UserById)

	log.Printf("Start Go HTTP Server")

	//port監視と実行
	err := http.ListenAndServe(":18081", router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
