package main

import "fmt"

func invertendoSinal(num int) int {
	return num * -1
}
func inverteSinalPonteiro(num *int) {
	*num = *num * -1
}
func main() {

	numero := 20
	fmt.Println(invertendoSinal(numero))
	fmt.Println(numero)

	num2 := 10
	fmt.Println(num2)
	inverteSinalPonteiro(&num2)
	fmt.Println(num2)
}
