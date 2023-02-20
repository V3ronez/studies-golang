package main

import "fmt"

func tentandoRecover() {
	if r := recover(); r != nil {
		fmt.Println("funcao recuperada")
	}
}

func aluno_aprovado(n1, n2 uint8) bool {
	defer tentandoRecover()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	}
	if media < 6 {
		return false
	}

	panic("nota 6, deu panic")

}
func main() {

	//panic - encerra o sistema mas antes chama todas as funcoes com defer

	fmt.Println("checando se aluno está aprovado")
	fmt.Println(aluno_aprovado(6, 6))
}
