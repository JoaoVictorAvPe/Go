package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var num int8

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		adiciona(&num)
		waitGroup.Done()
	}()

	go func() {
		subtrai(&num)
		waitGroup.Done()
	}()
	
	waitGroup.Wait()

	fmt.Println("RESULTADO:",num)
}

func adiciona(p *int8) {
	for i := 0; i < 5; i++ {
		*p += 1
		fmt.Println("ADICIONADO",*p)
		time.Sleep(time.Second)
	}
}

func subtrai(p *int8) {
	for i := 0; i < 5; i++ {
		*p += -1
		fmt.Println("SUBTRAIDO",*p)
		time.Sleep(time.Second)
	}
}