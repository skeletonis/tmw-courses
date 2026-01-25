package main

import "fmt"


func main(){

var choice uint
var total float64
var quantity uint

const agua_mineral float64 = 1.50
const agua_gas float64 = 2.5

fmt.Println("Qual garrafa de água você deseja?")
fmt.Println("1 - Água Mineral Natural: R$ 1.50 / 2 - Água Mineral com Gás: R$ 2.50")
fmt.Scanf("%d", &choice)

switch choice {
case 1:
	fmt.Println("Quantas garrafas de água você deseja?")
	fmt.Scanf("%d", &quantity)
	total = float64(quantity) * agua_mineral
	fmt.Printf("A sua compra deu %.2f reais", total)

case 2:
	fmt.Println("Quantas garrafas de água você deseja?")
	fmt.Scanf("%d", &quantity)
	total = float64(quantity) * agua_gas
	fmt.Printf("A sua compra deu %.2f reais", total)

default:
	fmt.Println("Escolhar uma opção válida!")
	break
}



}

