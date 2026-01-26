package main 

import ("fmt")

func square(a int) int{
	return a*a

}

func main(){
	
	var num int

	fmt.Printf("Descubra qual o quadrado de um número. Digite o número -> ")
	fmt.Scanf("%d", &num)

	square := square(num)

	fmt.Printf("O quadrado do número %d é %d", num, square)



}
