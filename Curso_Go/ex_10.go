package main

import "fmt"


func main(){

	var palavra string
	var count uint

	fmt.Println("Contador de letra a")
	fmt.Printf("Digite uma palavra que faremos a contagem da letra a -> ")
	fmt.Scanf("%s", &palavra)

	for _, a := range palavra{
		if (string(a) == "a"){
				count += 1
		}

	}
		
	fmt.Printf("A palavra %s possui %d letras a", palavra, count)

}


