package main

import (
	"fmt"
	"math"
	)

func main(){

var num float64

fmt.Println("Calculadora de raiz quadrada.")
fmt.Println("Digite qualquer número possível para a realizar sua matrícula hoje: ")
fmt.Scanf("%f", &num)

sqrt := math.Sqrt(num)
fmt.Printf("O valor da raiz quadrada é %.1f \n", sqrt)

}
