package moduls

import (
	"math/rand"
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
	movie_id int
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

func (m *Movie)Create() *Movie{
	m.ID = rand.Intn(100000000000000)
	res, err :=db.Query("INSERT INTO `movie` (`created_at`, `id`, `isbn`, `title`) VALUES (NOW(), ?, ?, ?)", m.ID, m.Isbn, m.Title)
	if err != nil{
		log.Fatal(err)
	}
	defer res.Close()
	d := m.Director
	d.movie_id = m.ID
	d.id = rand.Intn(100000000000000)
	res, err = db.Query("ISERT INTO `director`(`id`, `firstname`, `lastname`, `movie_id`) VALUES(?,?,?,?)", d.id, d.Firstname, d.Lastname, d.movie_id)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	return m
}
