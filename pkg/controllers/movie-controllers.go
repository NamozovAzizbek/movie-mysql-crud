package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/moduls"
)

var NewMovie moduls.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := moduls.GetMovies()
	fmt.Println(movies, "<------------nishon------------>")
	res, _ := json.Marshal(movies)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func GetMovieById(w http.ResponseWriter, r *http.Request) {
// 	req := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(req)
// 	if err != nil {
// 		fmt.Println("errror while parsing")
// 	}
// 	nemMovie := moduls.GetMovie(id)
// 	res, _ := json.Marshal(nemMovie)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }

// func CreateMovie(w http.ResponseWriter, r *http.Request){
// 	var newMovie moduls.Movie
// 	utils.ParseBody(r, NewMovie)
// 	newMovie.Create()
// 	res, _ := json.Marshal(newMovie)
// 	w.Header().Set("Content-Type", "pkglication/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(res)
// }
