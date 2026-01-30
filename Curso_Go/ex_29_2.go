package main

import "fmt"


//Listagem de Todas as frutas.
func getFrutas() map[string]float64 {
	frutas := map[string]float64{
		"maçã":    1.50,
		"banana":  2.75,
		"uva":     1.90,
		"pera":    1.25,
		"laranja": 0.65,
		"limão":   1.25,
		"goiaba":  2.15,
		"abacaxi": 3.20,
		"jaca":    5.80,
	}

	return frutas

}


//Retorna o valor da fruta desejada
func getFrutaValor(fruta string, lista_frutas map[string]float64) float64 {

	if valor, ok := lista_frutas[fruta]; !ok {
		fmt.Println("Fruta não encontrada! Verifique a digitação.")
		return 0.0
	} else {
		return valor
	}
}

func main() {

	var total float64
	listaFrutas := getFrutas()

	for {
		var inputText string
		fmt.Println("Que frutas deseja comprar?")
		fmt.Println(listaFrutas)
		fmt.Printf("Escolhar -> ")
		fmt.Scanf("%s", &inputText) //Aqui precisa ser um scanf pois precisa tipar o dado que esta entrando e
		//assim conseguir o breake no laço.

		if inputText == "" {
			break
		} else {
			valor := getFrutaValor(inputText, listaFrutas)
			total += valor
		}

	}

	fmt.Println("O valor total da compra é de: R$", total)

}
