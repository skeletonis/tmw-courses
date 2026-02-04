package main

import "fmt"

func fatorial(n int) int {

	calc := 1

	for i := n; i > 0; i-- {

		calc *= i

	}

	return calc

}

func main() {

	var inputNumber int
	fmt.Println("Descubra o fatorial de um número.")
	fmt.Printf("Entre com o valor -> ")
	fmt.Scan(&inputNumber)

	fmt.Printf("O valor de %d! é: %d", inputNumber, fatorial(inputNumber))

}
