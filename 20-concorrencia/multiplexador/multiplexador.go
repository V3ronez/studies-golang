package main

import (
	"fmt"
	"time"
)

func main() {
	//PADRAO MULTIPLEXADOR
	// junção de dois ou mais canal em um canal de saida apenas

	mensagem := multiplexar(escrever("ola mundo"), escrever("test"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-mensagem)
	}
}
func multiplexar(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case menssagem := <-canal1:
				canalSaida <- menssagem
			case menssagem := <-canal2:
				canalSaida <- menssagem
			}

		}
	}()
	return canalSaida
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
