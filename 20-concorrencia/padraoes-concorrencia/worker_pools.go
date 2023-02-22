package main

import "fmt"

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num-2) + fibonacci(num-1)
}

// tarefas <-chan =  apenas recebe valores
// resultados chan<- =  apenas envia valores

func workerPools(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func main() {
	tarefas := make(chan int, 30)
	resultados := make(chan int, 30)

	go workerPools(tarefas, resultados)
	go workerPools(tarefas, resultados)
	go workerPools(tarefas, resultados)
	go workerPools(tarefas, resultados)

	for i := 0; i < 30; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 30; i++ {
		result := <-resultados
		fmt.Println(result)
	}

}
