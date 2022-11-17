package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nanichang/bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStore(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}