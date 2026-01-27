package main

import "fmt"

func soma(a, b float64) float64{

	return a+b

}

func main(){

	var value1, value2 float64

	fmt.Printf("Calculadora que só soma. Digite o primeiro número")
	fmt.Scanf("%f", &value1)

	fmt.Printf("Digite o segundo número")
	fmt.Scanf("%f", &value2)


	total := soma(value1, value2)

	fmt.Printf("A soma de %.2f com %.2f é igual a %.2f", value1, value2, total)



}
