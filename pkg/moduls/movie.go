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

func GetAllMovie() []Movie {
	db := config.GetDB()
	res, err := db.Query("SELECT m.id, m.title, m.isbn, d.id, d.firstname, d.lastname FROM movie m INNER JOIN director d on m.id = d.movie_id")

	defer res.Close()

	if err != nil {
		log.Fatal("getall dagi select query bajarishda xatolik")
	}
	var (
		movies []Movie
		m Movie
		d Director
	)
	for res.Next() {
		err := res.Scan(&m.ID, &m.Isbn, &m.Title, &d.id, &d.Firstname, &d.Lastname)

		if err != nil {
			log.Fatal("getall dan go o'zgaruvchisiga o'zlashtrishda xatolik")
		}
		m.Director = &d
		movies = append(movies, m)
	}
	return movies
}
