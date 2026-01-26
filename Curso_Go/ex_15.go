package main

import (
	"fmt"
	"os"
)

func main(){

	tipoSorvete := map[string]float64 {
		"casquinha": 	1.5,
		"cascão": 		2.5,
		"cestinha": 	4.0,
	}

	saborSorvete := map[string]float64{
		"morango": 		0.0,
		"creme": 			0.0,
		"chocolate": 	0.0,
	}

	coberturaSorvete :=map[string]float64{
		"caramelo":		1.5,
		"morango":		1.5,
		"chocolate":	1.5,
	}

	itens:= map[string]map[string]float64{
		"tipos": 			tipoSorvete,
		"sabores":			saborSorvete,
		"coberturas":	coberturaSorvete,
	}


	var tipo, sabor, cobertura string
	var total float64

	fmt.Printf("Escolhar um tipo [casquinha/cascão/cestinha]")
	fmt.Scanf("%s", &tipo)

	if valor, ok := itens["tipos"][tipo]; !ok {
		fmt.Println("Tipo Inválido!")
		os.Exit(1)
	}else{
		total+=valor
	}
	
	fmt.Printf("Escolhar um sabor [morango/creme/chocolate]")
	fmt.Scanf("%s", &sabor)

	if valor, ok := itens["sabores"][sabor]; !ok {
		fmt.Println("Tipo Inválido!")
		os.Exit(1)
	}else{
		total+=valor
	}

	fmt.Printf("Escolhar uma cobertura [caramelo/morango/chocolate]")
	fmt.Scanf("%s", &cobertura)

	if valor, ok := itens["coberturas"][cobertura]; !ok {
		fmt.Println("Tipo Inválido!")
		os.Exit(1)
	}else{
		total+=valor
	}

fmt.Printf("Você escolheu um sorvete %s, sabor %s com cobertura de %s e o valor total do seu pedido ficou: %.2f", tipo, sabor, cobertura, total)

}
