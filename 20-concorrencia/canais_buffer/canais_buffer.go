package main

import "fmt"

func main() {
	canal := make(chan string, 2) // canal com buffer de 2, so vai dar deadlock quando estourar o tamanho
	canal <- "Ola mundo\n"
	canal <- "go é daora"
	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)

}
