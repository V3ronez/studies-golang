package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // canal pode enviar e receber dados do tipo string

	go escrever("ola mundo", canal)
	// <- canal. esperando que o canal receba um valor.

	// mensagem := <-canal
	// mensagem vai receber o valor de canal

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 3; i++ {
		// fmt.Println(text)
		canal <- texto // mandando um valor pra dentro do canal
		time.Sleep(time.Second)
	}
	close(canal) // WARNING, cuidado com deadlock!!!!!!
}
