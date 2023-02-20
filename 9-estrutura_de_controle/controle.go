package main

import "fmt"

func main() {

	numero := 10
	if numero <= 20 {
		fmt.Println("o numero é menor que 20")
	}

	if outro_numero := 10; outro_numero == numero { // if init so aparece nesse bloco
		fmt.Println("o outro numero é igual a 10")
	}

}
