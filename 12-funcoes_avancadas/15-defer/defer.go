package main

import "fmt"

func calcMedia(n1, n2 uint) bool {
	fmt.Println("caculando media")
	defer fmt.Println("media sendo calculada, retorno em breve") // retorno antes do primeiro RETURN
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	}
	return false
}

func main() {
	// defer -  adiar /  adia um pedaço de código ate quando possivel.
	// fn sem retorno: ultimo momento antes dela acabar.
	// fn com retorno: ate achar o primeiro return

	fmt.Println(calcMedia(4, 4))
}
