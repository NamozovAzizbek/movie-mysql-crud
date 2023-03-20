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
	movies := moduls.GetMovies()
	res, _ := json.Marshal(movies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	req := mux.Vars(r)["id"]
	id, err := strconv.Atoi(req)

	if err != nil {
		fmt.Println("errror while parsing")
		return
	}
	movie := moduls.GetMovie(id)
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	newMovie := &moduls.Movie{}
	utils.ParseBody(r, &newMovie)
	newMovie.Create()
	res, _ := json.Marshal(newMovie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request){
	req := mux.Vars(r)["id"]
	id, err := strconv.Atoi(req)
	if err != nil {
		fmt.Println("errror while parsing")
		return
	}
	movie := moduls.Delete(id)

	res, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}