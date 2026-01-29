//Append só funciona com slices

package main


import "fmt"


func main() {
	
	var retorno [2]string
	
	fmt.Println("Digite seu primeiro nome: ")
	fmt.Scan(&retorno[0])

	fmt.Println("Digite seu sobrenome: ")
	fmt.Scan(&retorno[1])

	fmt.Println("Seu nome completo é ", retorno[0], retorno[1])
	fmt.Println(retorno)

}


