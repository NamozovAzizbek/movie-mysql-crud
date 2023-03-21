package moduls

import (
	"log"
	"os"

	"github.com/NamozovAzizbek/movie-mysql-crud/pkg/config"
)

type Director struct {
	//id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Movie struct {
	ID         int    `json:"id"`
	Isbn       string `json:"isbn"`
	Title      string `json:"title"`
	DirectorId int
	Director   *Director
}

// func init() {
// 	config.Connect()
// }

var db = config.Connect()

func GetMovies() []Movie {
	movies := make([]Movie, 0)
	var m Movie
	var d Director
	row, err := db.Query("SELECT m.movieId, m.title, m.isbn, m.directorId, d.firstname, d.lastname FROM movie m INNER JOIN director d on m.directorId = d.id")

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer row.Close()

	for row.Next() {
		err := row.Scan(&m.ID, &m.Title, &m.Isbn, &m.DirectorId, &d.Firstname, &d.Lastname)

		if err != nil {
			log.Fatal(err)
			return nil
		}
		m.Director = &d
		movies = append(movies, m)
	}
	return movies
}

func GetMovie(id int) *Movie {
	res, err := db.Query("SELECT m.movieId, m.title, m.isbn, m.directorId ,d.firstname, d.lastname FROM movie m INNER JOIN director d on m.directorId = d.id where m.movieId = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	var (
		m Movie
		d Director
	)
	for res.Next() {
		err := res.Scan(&m.ID, &m.Title, &m.Isbn, &m.DirectorId, &d.Firstname, &d.Lastname)
		if err != nil {
			log.Fatal(err)
		}
		m.Director = &d
	}
	return &m
}

func (m *Movie) Create() *Movie {
	//director bor yoki yo'qligini tekshirish uchun
	row, err := db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
	var id int
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	defer row.Close()
	if id == 0 { //agar director mavjud bo'lmasa uni yaratamiz
		row, err := db.Query("INSERT INTO `director`(`firstname`, `lastname`) VALUES(?,?)", m.Director.Firstname, m.Director.Lastname)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer row.Close()
	}
	// director id ni olamiz
	row, err = db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	m.DirectorId = id
	res, err := db.Query("INSERT INTO `movie` (`created_at`, `isbn`, `title`, `directorId`) VALUES (NOW(), ?, ?, ?)", m.Isbn, m.Title, id)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	return m
}

func Delete(id int) *Movie {
	movie := GetMovie(id)
	row, err := db.Query("DELETE FROM movie WHERE movieId = ?", id)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer row.Close()
	return movie
}

func (m *Movie) Update(id int) *Movie {
	// get directorId
	directoId := 0
	row, err := db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
	if err != nil {
		log.Fatal(err)
	}
	for row.Next() {
		err = row.Scan(&directoId)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	defer row.Close()
	if directoId == 0 {
		//agar director mavjud bo'lmasa uni yaratamiz
		row, err := db.Query("INSERT INTO `director`(`firstname`, `lastname`) VALUES(?,?)", m.Director.Firstname, m.Director.Lastname)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer row.Close()
	}
	// director id ni olamiz
	row, err = db.Query("SELECT id FROM director WHERE lastname = ? and firstname = ?", m.Director.Lastname, m.Director.Firstname)
	for row.Next() {
		err = row.Scan(&directoId)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}

	// yangilash update
	row, err = db.Query("UPDATE movie SET title = ?, isbn = ?, directorId = ? where movieId = ?", m.Title, m.Isbn, directoId, id)
	if err != nil {
		return nil
	}
	defer row.Close()
	return m
}
