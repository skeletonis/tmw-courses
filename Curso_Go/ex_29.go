package main

import (
"fmt"
"os"
)

func main() {

	
var total float64

	frutas := map[string]float64{
		"maçã": 		1.50,
		"banana": 	2.75,
		"uva": 			1.90,
		"pera": 		1.25,
		"laranja": 	0.65,
		"limão": 		1.25,
		"goiaba": 	2.15,
		"abacaxi": 	3.20,
		"jaca": 		5.80,
	}


	for{
		var inputText string
		fmt.Println("Que frutas deseja comprar?")
		fmt.Println(frutas)
		fmt.Printf("Escolhar -> ")
		fmt.Scanf("%s", &inputText) //Aqui precisa ser um scanf pois precisa tipar o dado que esta entrando e 
																//assim conseguir o breake no laço.

		if inputText == ""{
			break
		} else if valor, ok := frutas[inputText]; !ok{
			fmt.Println("Tipo Inválido!")
			os.Exit(1)
		} else {
		total += valor
}

}

	fmt.Println("O valor total da compra é de: R$", total)

}
