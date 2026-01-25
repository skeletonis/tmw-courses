package main

import "fmt"


func main(){

var tipo_choice, sabor_choice, cobertura_choice, confirm uint
var tipo_descript, sabor, cobertura_descript string
var valor_tipo, valor_cobertura, total float64


fmt.Println("Monte seu sorvete")
fmt.Printf("1 - Casquinha: R$ 1,50 / 2 - Cascão: R$ 2,50 / 3 - Cestinha: R$4,00 ->")
fmt.Scanf("%d", &tipo_choice)
switch tipo_choice {
case 1:
	valor_tipo = 1.5
	tipo_descript = "Casquinha"

case 2:
	valor_tipo= 2.5
	tipo_descript = "Cascão"

case 3:
	valor_tipo = 4.0
	tipo_descript = "Cestinha"

default:
	fmt.Println("Escolha uma opção válida!")
	break
}

fmt.Printf("Escolha o seu sabor: 1 - morango / 2 - creme / 3 - chocolate ->")
fmt.Scanf("%d", &sabor_choice)
switch sabor_choice {
case 1:
	sabor = "Morango"

case 2:
	sabor = "Creme"

case 3:
	sabor = "Chocolate"

default:
	fmt.Println("Escolha uma opção válida!")
	break
}

fmt.Printf("Escolha a sua cobertura: 1 - Caramelo: R$ 1,50 / 2 - Morango: R$ 2,50 / 3 - Chocolate: R$4,00 / 4 - Sem Cobertura ->")
fmt.Scanf("%d", &cobertura_choice)
switch cobertura_choice {
case 1:
	valor_cobertura = 1.5
	cobertura_descript = "cobertura de Caramelo"

case 2:
	valor_cobertura = 2.5
	cobertura_descript = "cobertura de Morango"

case 3:
	valor_cobertura = 4.0
	cobertura_descript = "cobertura de Chocolate"

case 4:
	cobertura_descript = "sem cobertura"

default:
	fmt.Println("Escolha uma opção válida!")
	break
}

total = valor_cobertura+valor_tipo
fmt.Printf("O seu pedido foi um %s - %0.2f - de sabor %s com %s - %.2f \n", tipo_descript, valor_tipo, sabor, cobertura_descript, valor_cobertura)
fmt.Printf("O total do seu pedido foi de R$ %.2f. \n", total)
fmt.Println(" Aperte: 1 - Confirmar pedido e 2 - Cancelar pedido")
fmt.Scanf("%d", &confirm)
if(confirm == 1){
	fmt.Println("Prossiga para o pagamento.")
} else {
	fmt.Println("Seu pedido será cancelado.")
}


}

