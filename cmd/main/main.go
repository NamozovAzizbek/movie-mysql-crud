package main

import (
	"log"
	"net/http"

	routers "github.com/NamozovAzizbek/movie-mysql-crud/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routers.RegisterMovieRouters(r)
	r.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}