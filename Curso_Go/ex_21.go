package main	

import "fmt"

func saudacao(a string){
	fmt.Printf("Olá %s! Boas Vinda!", a)
}


func main(){
	
	var input string

	fmt.Println("Olá, qual é o seu nome?")
	fmt.Scanf("%s", &input)

	saudacao(input)

}
