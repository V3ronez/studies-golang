package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // import implicito, quem vai esta usando é o pacote "database/sql" de moto implicito
)

func main() {

	var stringConeccao = "user=veronez password=261602317 dbname=golang_studies sslmode=disable"
	db, err := sql.Open("postgres", stringConeccao)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil { // comum reaproveitando a variavel de erro que está como <nil>
		log.Fatal(err)
	}

	linhas, err := db.Query("select from usuarios where 1=1")
	if err != nil {
		log.Fatal(err)
	}
	defer linhas.Close()
	fmt.Println(linhas)

	fmt.Println("deu certo amigo!")

}
