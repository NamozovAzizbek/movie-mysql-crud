package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/moduls"
	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/utils"
	"github.com/gorilla/mux"
)

var NewMovie moduls.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	newMovies := moduls.GetMovies()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	req := mux.Vars(r)["id"]
	id, err := strconv.Atoi(req)
	if err != nil {
		fmt.Println("errror while parsing")
	}
	nemMovie := moduls.GetMovie(id)
	res, _ := json.Marshal(nemMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	var movie moduls.Movie
	utils.ParseBody(r, movie)
}
