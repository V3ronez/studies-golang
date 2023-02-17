package main

import "fmt"

type user struct {
	name    string
	age     int8
	address address
}

type address struct {
	number     int8
	street     string
	complement string
}

func main() {

	usuario := user{"henrique", 20, address{number: 10, street: "rua do bobo"}}

	var usuario2 user
	usuario2.name = "lourenco"
	usuario2.age = 53

	usuario3 := user{name: "mendes"} //atualizando apenas um elemento
	fmt.Println(usuario, usuario2, usuario3)

}
