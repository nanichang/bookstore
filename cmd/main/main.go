package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com.jinhzu/gorm/dialects/mysql"
	"github.com/nanichang/bookstore/pgk/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}