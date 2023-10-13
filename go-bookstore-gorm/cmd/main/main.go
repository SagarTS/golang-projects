package main

import (
	"go-bookstore/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9000",r))
}