package main

import (
	"fmt"
	"time"
)

func main() {
	//goroutines
	go escrever("escrevendo daora 1")
	escrever("escrevendo daora 2")
}

func escrever(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
