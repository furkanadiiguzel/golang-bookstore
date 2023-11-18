package main

import (
	"log"
	"net/http"

	"github.com/furkanadiiguzel/golang-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))

}
