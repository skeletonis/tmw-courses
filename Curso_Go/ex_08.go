package main 

import (
				"fmt"
				"strings"
)

func main() {
	
var name string
fmt.Println("Qual o seu nome: ")
fmt.Scanf("%s", &name)

if strings.Contains(name, "Lima") {
	fmt.Println("Você faz parte da família Lima")
} else{
	fmt.Println("Você não faz parte da família Lima")
}

}
