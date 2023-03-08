package main

func main() {
	// PONTEIROS RECEBEM ESPACOS DA MEMORIA
	//ponteiros sao referencias de memorias e não guardam o valor em si da variavel mas sim a referencia dela na memoria
	// * indica que é um ponteiro e & voce referencia outra variavel
	var num1 int = 12
	var ponteiro_num1 *int = &num1 // ponteiro_num1 tem o valor da referencia na memoria de num1
	println(num1, ponteiro_num1)
	num1 = 30
	println(num1, *ponteiro_num1) // * ler o valor que esta naquele espaco de memoria (desreferenciar)

}
