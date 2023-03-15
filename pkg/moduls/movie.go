package moduls

import (
	"log"

	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/config"
)

type Movie struct {
	ID       int    `json:"id"`
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	Director *Director
}

type Director struct {
	id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func init() {
	config.Connect()
}

var db = config.GetDB()

func GetMovies() []Movie {
	res, err := db.Query("SELECT m.id, m.title, m.isbn, d.id, d.firstname, d.lastname FROM movie m INNER JOIN director d on m.id = d.movie_id")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	var (
		movies []Movie
		m      Movie
		d      Director
	)
	for res.Next() {
		err := res.Scan(&m.ID, &m.Isbn, &m.Title, &d.id, &d.Firstname, &d.Lastname)

		if err != nil {
			log.Fatal(err)
		}
		m.Director = &d
		movies = append(movies, m)
	}
	return movies
}
func GetMovie(id int) *Movie {
	res, err := db.Query("SELECT m.id, m.title, m.isbn, d.id, d.firstname, d.lastname FROM movie m INNER JOIN director d on m.id = d.movie_id where m.id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	var (
		m Movie
		d Director
	)
	for res.Next() {
		err := res.Scan(&m.ID, &m.Isbn, &m.Title, &d.id, &d.Firstname, &d.Lastname)
		if err != nil {
			log.Fatal(err)
		}
		m.Director = &d
	}
	return &m
}

func (m *Movie)Create() {
	
}
