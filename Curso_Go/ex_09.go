package main

import "fmt"

func main() {

var note1, note2, note3, media float64

fmt.Printf("Nota da sua primeira avaliação -> ")
fmt.Scanf("%f", &note1)

fmt.Printf("Nota da sua segunda avaliação -> ")
fmt.Scanf("%f", &note2)

fmt.Printf("Nota da sua terceira avaliação -> ")
fmt.Scanf("%f", &note3)

media = (note1+note2+note3) / 3

if(media >= 6){
	fmt.Printf("Parabéns, você passou com média %.2f", media)
} else if(media < 6 && media >= 4){
	fmt.Printf("Você ficou de recuperação, sua média é de  %.2f", media)
} else {
	fmt.Printf("Que pena você reprovou! Sua média é %.2f", media)
}

}
