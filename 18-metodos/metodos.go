package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvarUsuario() {
	fmt.Printf("o usuario %s foi salvo com sucess!\n", u.nome)
}
func (u *usuario) fazerAniversario() { // usar o ponteiro pra modificar os atributos de um struct
	u.idade++
}

func main() {
	usuario1 := usuario{nome: "geo", idade: 20}
	fmt.Println(usuario1)
	usuario1.salvarUsuario()
	usuario1.fazerAniversario()
	fmt.Println(usuario1)
}
