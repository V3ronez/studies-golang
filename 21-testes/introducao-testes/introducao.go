package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoEndereco("rua dos bobos")
	fmt.Println(tipoEndereco)
}
