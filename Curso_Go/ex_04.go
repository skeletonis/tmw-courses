package main

import "fmt
	math"

func main(){

var num float64

fmt.Printf("Calculadora de raiz quadrada. Digite qualquer número possível para a realizar sua matrícula hoje: \n")
fmt.Scanf("%d", &num)

sqt := math.Sqrt(num)
fmt.Printf("O valor da raiz quadrada é ", sqt)

}
