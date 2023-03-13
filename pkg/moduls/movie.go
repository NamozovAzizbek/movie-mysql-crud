package moduls

import "github.com/NamozovAzizbek/movie-mysql-crud/pkg/config"

type Movie struct {
	ID       int    `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	Director *Director
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func init(){
	config.Connect()

}
