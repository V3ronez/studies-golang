package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("escrevendo do main")
	auxiliar.Escrevendo()

	error := checkmail.ValidateFormat("v3ronez.dev@gmail.com")
	fmt.Println(error)
}
