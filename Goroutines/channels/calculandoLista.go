package main

import (
	"fmt"
	"time"
)

func main() {
	listChannel := make(chan []int, 2)
	resultChannel := make(chan int, 2)
	go calculatePairs(listChannel, resultChannel)
	go calculateNotPairs(listChannel, resultChannel)

	const arraySize int = 5
	var numberList []int
	for i := 0; i < arraySize; i++ {
		numberList = append(numberList, i+1)
	}
	fmt.Println(numberList)
	


	listChannel <- numberList
	listChannel <- numberList


	result := <-resultChannel
	fmt.Println("RESULTADO RETORNADO NO CANAL",result)
	result2 := <-resultChannel
	fmt.Println("RESULTADO RETORNADO NO CANAL",result2)

}

func calculatePairs(listChannel <-chan []int, resultChannel chan<- int) {

	numbersList := <- listChannel
	fmt.Println("PARES, LISTA RECEBIDA", numbersList)

	var sum int
	for _,number := range numbersList {
		isPairNumber := number % 2 == 0
		if isPairNumber {
			sum += number
		}
	}
	fmt.Println("PARES - VALOR ENVIADO PARA O CANAL:",sum)
	resultChannel <- sum
}

func calculateNotPairs(listChannel <-chan []int, resultChannel chan<- int) {

	numbersList := <- listChannel
	fmt.Println("IMPARES, LISTA RECEBIDA", numbersList)

	var sum int
	for _,number := range numbersList {
		isPairNumber := number % 2 != 0
		if isPairNumber {
			sum += number
		}
		time.Sleep(time.Second)
	}
	fmt.Println("IMPARES - VALOR ENVIADO PARA O CANAL:",sum)
	resultChannel <- sum
}

