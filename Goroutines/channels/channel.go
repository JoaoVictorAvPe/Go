package main

import (
	"fmt"
	"time"
)

func main1() {
	channel := make(chan map[string]int8)
	channel2 := make(chan map[string]int8)

	go corredor(1, channel)
	go corredor(2, channel2)

	vencedor := <- channel
	vencedor2 := <- channel2

	fmt.Println("VENCEDOR:",vencedor["ID"])
	fmt.Println("PONTOS:",vencedor["PONTOS"])

	fmt.Println("VENCEDOR:",vencedor2["ID"])
	fmt.Println("PONTOS:",vencedor2["PONTOS"])
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