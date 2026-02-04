package main

import "fmt"

func fatorial(n int) int {
	if n == 0{
		return 1
	} else {
		return n * fatorial(n-1)
	}
}

func main() {

	var inputNumber int
	fmt.Println("Descubra o fatorial de um número.")
	fmt.Printf("Entre com o valor -> ")
	fmt.Scan(&inputNumber)

	fmt.Printf("O valor de %d! é: %d", inputNumber, fatorial(inputNumber))

}
