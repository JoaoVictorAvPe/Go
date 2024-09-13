package main

import "fmt"

func main() {
	var soma, _ int8 = calcular(2, 4)
	fmt.Println(soma)

	var f = func() string {
		var texto string = "Função F"
		fmt.Println(texto)
		return texto
	}

	var retorno string = f()
	fmt.Println(retorno)
}

func calcular(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}
