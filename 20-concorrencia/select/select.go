package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "mensagem canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 3)
			canal2 <- "canal dois recebendo mensagem "
		}
	}()

	for {
		// mensagem1 := <-canal1
		// fmt.Println(mensagem1)

		// mensagem2 := <-canal2
		// fmt.Println(mensagem2)

		// está limitando a msg1 a aparecer na tela, ela so aparece apos a msg2 receber o dado ( mesmo msg 1 sendo mais rapida )

		select { // assim que o canal 1 manda a msg ela já é printada na tela, sem precisar esperar o canal 2 que é mais lento
		case mensagem1 := <-canal1:
			fmt.Println(mensagem1)
		case mensagem2 := <-canal2:
			fmt.Println(mensagem2)

		}

	}
}
