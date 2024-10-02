package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan map[string]int8)

	go corredor(1, channel)
	go corredor(2, channel)

	vencedor := <- channel
	fmt.Println("VENCEDOR:",vencedor["ID"])
	fmt.Println("PONTOS:",vencedor["PONTOS"])
}

func corredor(id int8, channel chan map[string]int8) {
	var pontos int8
	for i := 0; i < 10 ; i++ {
		pontos += 1
		fmt.Println("CORREDOR", id, "NA FRENTE")
		time.Sleep(time.Second)
	}
	data := map[string]int8{"ID":id, "PONTOS":pontos}
	channel <- data
}