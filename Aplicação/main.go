package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	aplication := app.Gerar()
	if error := aplication.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}