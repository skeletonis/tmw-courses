package main	

import (
	"fmt"
	"math"
)

func calc(r, s float64)float64{

	return math.Pow(r, s)


}


func main(){

	var input1, input2 float64
	fmt.Printf("Descubra pontênciações. Digite o número que servirá como base -> ")
	fmt.Scanf("%f", &input1)

	fmt.Printf("Digite o número que será o expoente -> ")
	fmt.Scanf("%f", &input2)


	resultado := calc(input1, input2)

	fmt.Printf("O número de base %.2f, elevado a %.2f é igual a %.2f", input1, input2, resultado)


}
