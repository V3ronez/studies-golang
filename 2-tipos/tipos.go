package main

import "errors"

func main() {
	// var numero1 uint8 = -123 // erro
	var numero2 int8 = 23
	var numero3 int = 1000 //por padrao puxa a base do seu pc  (nesse caso int64)
	var numero4 float32 = 13.23
	var numero5 float64 = 13.5353535353535353

	// conceita do valor 0
	//valor attr a uma var quando voce não inicializa ela, pra string é "" e pra int é o valor 0, pra error é nil
	var err error = errors.New("erro interno")

	println(numero2, numero3, numero4, numero5, err)
}
