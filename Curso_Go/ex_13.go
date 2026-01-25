package main 

import (
	"fmt"
	"slices"
)

func main(){
	var notas [4]float64
	var summ float64

	for i := range notas{
		var nota float64
		fmt.Printf("Digite a sua nota %d: ", (i+1))
		fmt.Scanf("%f", &nota)
		notas[i] = nota
				
	}

	min := slices.Min(notas[:])
	max := slices.Max(notas[:])

	for _, value := range notas{
		summ += value
	}

fmt.Println("Média: ", summ/float64(len(notas)))
fmt.Println("Máxima: ", max)
fmt.Println("Mínima: ", min)

}
