package main

import "fmt"

func saudacao(a string) {
	fmt.Printf("Olá %s! Boas Vindas!", a)
}

func main(){
	var name string
	fmt.Println("Qual o seu nome?")
	fmt.Scanf("%s", &name)
	saudacao(name)
}
