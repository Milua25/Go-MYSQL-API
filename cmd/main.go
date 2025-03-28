package main

import (
	"fmt"
	"github.com/Golang-Personal-Projects/Go-Projects/03-Go-Mysql-API/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterBookStoreRoutes(r)

	http.Handle("/", r)
	fmt.Println("Starting Server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
