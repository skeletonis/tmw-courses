package main

import "fmt"

func primo(n int) bool {
	if n%1 == 0 && n%n == 0 {
		if n%2 != 0 || n==2 && n%3 != 0  || n==3  && n%5 != 0  || n==5  {
			return true
		}}
		return false

}

func main() {

	var inputNumber int

	fmt.Println("Verifique se o numero e primo")
	fmt.Printf("Digiter o numero -> ")
	fmt.Scan(&inputNumber)

	if primo(inputNumber) {
		fmt.Printf("E primo")
	} else {

		fmt.Printf("Nao e primo")
	}

}
