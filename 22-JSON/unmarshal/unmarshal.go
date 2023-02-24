package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome      string `json:"nome"` // `json:"-"` omite essa chave quando parsear pra JSON
	Sobrenome string `json:"sobrenome"`
	Idade     uint   `json:"idade"`
}

func main() {

	cachorroEmJSON := `{"nome":"dudu","sobrenome":"ruim","idade":4}`
	var c cachorro

	dogEmByte := []byte(cachorroEmJSON) // transformando um json em um slice de byte code
	if erro := json.Unmarshal(dogEmByte, &c); erro != nil {
		log.Fatal(erro.Error())
	}
	fmt.Println(c) // {dudu ruim 4}
	fmt.Println("---------------------------")

	// enderecoEmMap := map[string]string{
	// 	"rua":    "do bobo",
	// 	"numero": "2",
	// 	"bairro": "centro",
	// }

	dogEmJSON := `{"nome":"toni","sobrenome":"da silva"}`

	dog := make(map[string]string) //map[nome:toni sobrenome:da silva]

	if err := json.Unmarshal([]byte(dogEmJSON), &dog); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(dog) // {dudu ruim 4}
}
