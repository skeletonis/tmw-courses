package main

import "fmt"

func imparPar(numero int){
	
	if numero % 2 == 0 {
		fmt.Printf("O número %d é par.", numero)
	} else {
		fmt.Printf("O número %d é ímpar.", numero)
	}


}

func main() {
	var input int

	fmt.Println("Teste impar ou par.")
	fmt.Printf("Digite um número e descubra se ele é impar ou par -> ")
	fmt.Scanf("%d", &input)


	imparPar(input)

}
