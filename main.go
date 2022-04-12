package main

import (
	"fmt"
	"go-web-template/server"
	"go-web-template/server/config"

	"github.com/gorilla/mux"
)

func main() {
	run()
}

func run() {
	fmt.Println("Starting...")
	db := config.CreateConnection()
	fmt.Println("connect DB success")

	router := mux.NewRouter()
	port := ":9999"

	server.StartServer(router, port, db)
}
