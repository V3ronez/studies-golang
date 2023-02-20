package main

import "fmt"

func main() {

	// o go so tem o loop for
	// structs não podem ser interaveis

	// i := 0
	// for i < 10 {
	// 	i++
	// 	time.Sleep(time.Second)
	// 	fmt.Println("for do i")
	// }

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("for do j")
	// }

	// array_interavel := [3]string{"joao", "carlos", "ana"}
	// for key, value := range array_interavel {
	// 	fmt.Println(key, value)
	// }

	// for _, value := range array_interavel {
	// 	fmt.Println(value)
	// }

	slice1 := []string{"carro", "moto", "bike", "navio"}

	for _, value := range slice1 {
		fmt.Println(value)
	}
}
