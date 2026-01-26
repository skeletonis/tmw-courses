package main	

import "fmt"

func inverter(a, b int) (int, int){
	c, d := b, a

	return c, d
}

func main(){

var a, b int

fmt.Printf("Digite o valor de a -> ")
fmt.Scanf("%d", &a)

fmt.Printf("Digite o valor de b -> ")
fmt.Scanf("%d", &b)

n_a, n_b := inverter(a, b)

fmt.Printf("Os valores invertidos ficaram como a = %d e b = %d", n_a, n_b)

}
