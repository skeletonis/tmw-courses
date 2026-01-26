package main	

import (
	"fmt"
	"strconv"
)

func media(n []float64) float64{
	var sum, media float64

	for _, values := range n{
		sum += values
	}

	media = sum/float64(len(n))

	return media

}


func main(){

	var numeros = []float64{}

	for{
	var input string
	fmt.Printf("Entre com o valor: ")
	fmt.Scanf("%s", &input)

	if input == ""{
		fmt.Println("A media dos números fornecidos é de ", media(numeros))
		break
	} 

	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Entre com um valor válido.")
	}

	numeros = append(numeros, value)


	}


}
