package main

func main() {
	//ponteiros sao referencias de memorias e não gurdam o valor em si da variavel mas sima  referencia dela na memoria

	var num1 int = 12
	var ponteiro_num1 *int = &num1 // * indica que é um ponteiro e & voce referencia outra variavel
	println(num1, ponteiro_num1)
	num1 = 30
	println(num1, *ponteiro_num1) // * para pegar o valor do ponteiro em vez da referencia (desreferenciar)

}
