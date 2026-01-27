package main	

import (
	"fmt"
	"math"
)

func calc(r float64)(float64, float64){

	area := math.Pi*math.Pow(r, 2)
	perimeter := 2*(r*math.Pi)

	return area, perimeter

}


func main(){

	var input float64
	fmt.Printf("Descubra as dimensões de uma circunferência. Digite o raio -> ")
	fmt.Scanf("%f", &input)


	area, perimeter := calc(input)

	fmt.Printf("A área da circunferência é: %.2f. E o seu perímetro é: %.2f", area, perimeter)


}
