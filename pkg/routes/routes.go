package routers

import(
	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterMovieRouters = func (router *mux.Router)  {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.GetMovieById).Methods("GET")
	router.HandleFunc("/movie", controllers.CreateMovie).Methods("POST")
	// router.HandleFunc("/movies{id}", controllers.UpdatMovie).Methods("PUT")	
	router.HandleFunc("/movie/{id}", controllers.DeleteMovie).Methods("DELETE")
}