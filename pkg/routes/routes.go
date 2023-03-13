package routers

import(
	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieRouters = func (router *mux.Router)  {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movies", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies", controllers.UpdatMovie).Methods("PUT")	
	router.HandleFunc("/movies", controllers.DeleteMovies).Methods("DELETE")
}