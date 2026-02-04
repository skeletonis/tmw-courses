package main	

import "fmt"

func multi(n1, n2 int)int{
	return n1*n2
}

func main(){

	var inputNumber, calc int

	fmt.Println ( "Descubra a tabuada de um numero" )
	fmt.Printf("Digite um numero -> ")
	fmt.Scan(&inputNumber)

	for i:=1; i <= 10; i++{
		calc = multi(inputNumber, i)
		fmt.Printf("%d x %d = %d \n", inputNumber, i, calc)

	}


}
