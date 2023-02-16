package main

import "fmt"

func calcs(num1 int16, num2 int16) (int16, int16) {
	soma := num1 + num2
	subtracao := num1 - num2
	return soma, subtracao
}
func somando(num1, num2 int) int {
	return num1 + num2
}

func main() {
	soma, _ := calcs(17, 19) // pegando apenas um retorno da funcao

	somandoValores := somando(101, 231)
	var mostraNaTela = func(text string) {
		println(text)
	}
	mostraNaTela("test")
	fmt.Println(soma, somandoValores)
}
