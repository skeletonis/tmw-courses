package main

import "fmt"

func main(){

var height, summ float64

fmt.Println("Programa de soma de altura")

for i := 1; i<=4; i++ {
	fmt.Printf("Entre com a altura %d -> ", i)
	fmt.Scanf("%f", &height)
	summ += height

}

fmt.Printf("O somatório de todas as alturas é igual %.2f", summ)

}
