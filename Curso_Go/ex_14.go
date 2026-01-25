package main 

import "fmt"

func main(){
	var temp [6]float64
	var summ float64

	for i := range temp{
		var t float64
		fmt.Printf("Digite a temperatura %d: ", (i+1))
		fmt.Scanf("%f", &t)
		temp[i] = t
				
	}


	for i:=1; i<len(temp)-1; i++{
		summ += temp[i]
	}

fmt.Println(summ)
media:= summ/float64(len(temp)-2)

if media > 30 {
	fmt.Println("Houve um aumento de temperatura!")
	fmt.Printf("A media de temperatura atual é: %.2f", media)
} else{
	fmt.Println("Temperatura normal")
	fmt.Printf("A media de temperatura atual é: %.2f", media)
}

}
