package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("teste de canal") //obs: a seta (<-escrever("teste de canal")) pode ser usada pra retornar o valor de uma funcao que retorna um canal
	for i := 0; i < 10; i++ {
		fmt.Println(<-canal) //obs: sem a seta (<-canal), retorna a referencia da memoria.
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return canal
}
