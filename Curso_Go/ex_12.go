package main 

import (
			"fmt"
			"strconv"
		)

func main(){

	var summ float64


	for{
	
		var input string
		fmt.Printf("Entre com o valor -> ")
		fmt.Scanf("%s", &input)


		if(input == ""){
			break
		} else {
		
			valor, erro := strconv.ParseFloat(input, 64)
		

			// Quando algo da errado -> ela contém um valor e quando tudo dá certo o erro da -> nil que é o equivalente de nulo,
			// ou seja, nada deu errado. Nesse caso, a condição quer dizer que - se erro, faça isso.
			if erro != nil{
				fmt.Println("Entre com um valor válido!")
				break
			}
			summ += valor	

		}

	}

		fmt.Printf("Saldo total em conta: R$ %.2f", summ)

}



