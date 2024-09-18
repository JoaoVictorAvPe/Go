package main

import (
	"fmt"
	"module/people"
)

func main() {
	pessoa1 := people.People{Name: "Joao", Age: 23, Gender: "Male"}
	fmt.Println(pessoa1.GetName())

	pessoa1.SetName("Maria")

	fmt.Println(pessoa1.GetName())

	nums := []int{3, 4, 5, 6}

	nums = append(nums, 4)

	for  i := 0; i<len(nums); i++ {
		fmt.Println(nums[i])
	}
}
