package main

import (
	"net/http"
	"log"
	"github.com/cappyzawa/go-tiny-crud/api"
)

func main() {
	//routingはここ
	http.HandleFunc("/", api.Hello)
	//この下にrouting定義を増やしていけばいい感じな気がする。

	log.Printf("Start Go HTTP Server")

	//port監視と実行
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
