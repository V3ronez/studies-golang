package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// wait group - sincronizar as funcoes concorrentes
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("escrevendo...")
		waitGroup.Done()
	}()
	go func() {
		escrever("lendo...")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func escrever(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
