package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type pessoa struct {
	Name 		string	`json:"name"`
	LastName	string	`json:"lastName"`
	Weigth 		float32	`json:"weigth"`
}

func (p pessoa) toJSON() *bytes.Buffer {
	pBytesJSON, error := json.Marshal(p)
	if error != nil {
		log.Fatal(error)
	}

	return bytes.NewBuffer(pBytesJSON)
}

func fromJSONtoStruct(pessoaJSON string) pessoa {
	var p1 pessoa
	if erro := json.Unmarshal([]byte(pessoaJSON), &p1); erro != nil {
		log.Fatal(erro)
	}

	return p1
}

func main() {
	var p pessoa = pessoa{"Joao Victor", "Avila", 78.9}

	fmt.Println(p.toJSON())

	var p1 string = `{"name":"Joao Victor","lastName":"Avila","weigth":80.9}`

	pessoa1 := fromJSONtoStruct(p1)
	fmt.Println(pessoa1.toJSON())
}