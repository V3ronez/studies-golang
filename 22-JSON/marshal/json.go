package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Nome      string `json:"nome"`
	Sobrenome string `json:"sobrenome"`
	Idade     uint   `json:"idade"`
}

func main() {
	p := pessoa{Nome: "veronez", Sobrenome: "bobo", Idade: 23}

	// fmt.Println(p)
	pessoaJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(pessoaJSON)
	fmt.Println(bytes.NewBuffer(pessoaJSON))
	fmt.Println("-----------------------------")

	cachorro := map[string]string{
		"nome": "dudu",
		"raca": "ruim",
	}
	c, erro := json.Marshal(cachorro)
	if erro != nil {
		log.Fatal(erro.Error())
	}
	fmt.Println(c)
	fmt.Println(bytes.NewBuffer(c))
}
