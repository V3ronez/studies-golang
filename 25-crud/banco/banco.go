package banco

import (
	"database/sql"

	_ "github.com/lib/pq" // import implicido que o sql do golang precisa
)

func Connectar() (*sql.DB, error) {
	strConnect := "user=veronez password=261602317 dbname=golang_studies sslmode=disable"
	db, err := sql.Open("postgres", strConnect)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
