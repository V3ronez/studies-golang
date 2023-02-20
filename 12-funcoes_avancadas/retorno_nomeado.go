package main

import "fmt"

func calcs(n1, n2 int) (soma int, suntracao int) {
	soma = n1 + n2
	suntracao = n1 - n2
	return
}

func main() {

	sum, sub := calcs(10, 20)
	fmt.Println(sum, sub)
}
