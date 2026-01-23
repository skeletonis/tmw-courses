package main

import "fmt"

func main(){
	
	var name string

	fmt.Println("Bom dia!")
	fmt.Printf("Qual o seu nome? ")
	fmt.Scanf("%s", &name)

	fmt.Printf("Prazer em te conhecer %s!\n", name)


}
