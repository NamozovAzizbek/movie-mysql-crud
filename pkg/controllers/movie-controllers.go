package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/moduls"
)

var NewMovie moduls.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	newMovies := moduls.GetAllMovie()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
