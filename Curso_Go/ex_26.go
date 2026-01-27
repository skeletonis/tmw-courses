package main

import "fmt"

func converte(totalSegundos int)(h, m, s int) {
	h = totalSegundos/ 3600
	resto := totalSegundos % 3600

	m = resto / 60
	s = resto % 60

	return h, m, s
}
	


func main() {
	var input int
	fmt.Println("Conversor de segundos para horas, minutos e segundos.")
	fmt.Printf("Entre com o valor em segundos: ")
	fmt.Scanf("%d", &input)

	h, m, s := converte(input)

	fmt.Printf("%d:%d:%d\n", h, m, s)
}
