package main

import "fmt"


func main(){

var num int

fmt.Printf("Digite um número para saber qual o dobro: ")
fmt.Scanf("%d", &num)

double := num * 2 
fmt.Printf("O dobro de %d é igual a %d", num, double)

}

