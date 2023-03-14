package controllers

import (
	"net/http"

	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/moduls"
)

var NewMovie moduls.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	newMovie := moduls.GetAllMovie()

}
