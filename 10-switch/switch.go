package main

import "fmt"

func DiaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "domingo"
		// fallthrough faz cair dentro do segundo case direto
	case 2:
		return "segunda"
	default:
		return "nenhum dia foi passado"
	}
}

func main() {
	fmt.Println(DiaDaSemana(9))
}
