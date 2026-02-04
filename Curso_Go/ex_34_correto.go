package main

import "fmt"

func primo(n int) bool {
	if n == 1{
		return true
	} else if n = 2 {
		return true
	} else if n % 2 == 0 {
		return false
	}

	// Ja testamos as outras possibilidades anteriormente por isso ja comeca com 3;
	// i*i<=n verifica todos os numeros ate que o quadrado seja igual ao valor de n, pois -
	// - se um número n não tem nenhum divisor até a raiz quadrada dele, então ele não tem divisor nenhum.
	// dessa forma, entramos na definicao de que o numero primo so pode ser divido por 1 e por ele mesmo.
	for i := 3; i*i <= n; i+=2 {
		if n % i == 0{
			return false
		}
	}

return true
}

func main() {

	var inputNumber int

	fmt.Println("Verifique se o numero e primo")
	fmt.Printf("Digiter o numero -> ")
	fmt.Scan(&inputNumber)

	if primo(inputNumber) {
		fmt.Printf("E primo")
	} else {

		fmt.Printf("Nao e primo")
	}

}
