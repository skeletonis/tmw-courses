package main

import (
"fmt"
"os"
)


func getValor(fruta string) float64{

var valor float64

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

	if preco, ok := frutas[fruta]; !ok{
		fmt.Println("Tipo inválido!")
		os.Exit(1)

	} else {
		valor = preco 
	}
	
return valor

}


func main() {

var total float64

	for{
		var inputText string
		fmt.Println("Que frutas deseja comprar?")
		fmt.Println("maçã / banana / uva / pera / laranja / limão / goiaba / abacaxi / jaca")
		fmt.Printf("Escolhar -> ")
		fmt.Scan(&inputText)

		if inputText == ""{
			fmt.Printf("O valor total da compra é de: R$", total)
			break
	}
		
		total += getValor(inputText)



}
}
