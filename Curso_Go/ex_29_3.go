package main

import "fmt"

type Frutas struct {
	frutas map[string]float64
}

func (f *Frutas) getFrutas() map[string]float64 {
	return f.frutas
}

func (f *Frutas) getFrutaValor(fruta string) float64 {

	if valor, ok := f.frutas[fruta]; !ok {
		fmt.Println("Fruta não encontrada! Verifique a digitação.")
		return 0.0
	} else {
		return valor
	}
}

func main() {

	var total float64
	listaFrutas := Frutas{
		frutas: map[string]float64{
			"maçã":    1.50,
			"banana":  2.75,
			"uva":     1.90,
			"pera":    1.25,
			"laranja": 0.65,
			"limão":   1.25,
			"goiaba":  2.15,
			"abacaxi": 3.20,
			"jaca":    5.80,
		},
	}

	for {
		var inputText string
		fmt.Println("Que frutas deseja comprar?")
		fmt.Println(listaFrutas.getFrutas())
		fmt.Printf("Escolha -> ")
		fmt.Scanln(&inputText)

		if inputText == "" {
			break
		} else {
			valor := listaFrutas.getFrutaValor(inputText)
			total += valor
		}

	}

	fmt.Println("O valor total da compra é de: R$", total)

}
