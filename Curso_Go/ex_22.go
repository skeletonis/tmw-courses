package main	

import "fmt"

func saudacao(a string, b uint8){
	fmt.Printf("Olá %s, bom saber que você tem %d anos. Boas Vinda!", a, b)
}


func main(){
	
	var inputName string
	var inputAge uint8
		

	fmt.Println("Olá, qual é o seu nome?")
	fmt.Scanf("%s", &inputName)

	fmt.Println("Olá, qual é a sua idade?")
	fmt.Scanf("%d", &inputAge)



	saudacao(inputName, inputAge)

}
