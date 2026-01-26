package main

import "fmt"

func EhPar(a int) bool{

	if(a % 2 == 0){
		return true
	} else{ 
	return false}

}

func main() {

	var num int

	fmt.Printf("Descubra agora se um numero é par. Escolha um número -> ")
	fmt.Scanf("%d", &num)

	teste := EhPar(num)

	if teste{
		fmt.Printf("%d é um numéro par.", num)
	} else {
		fmt.Printf("%d é um numéro ímpar.", num)
	}
	



}
