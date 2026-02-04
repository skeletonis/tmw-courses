package main

import "fmt"

func FizzBuzz(n int) string {
	var text string

	if n%3 == 0 && n%5 == 0 {
		text = "FizzBuzz"
	} else {

		if n%3 == 0 {
			text = "Fizz"
		} else if n%5 == 0 {
			text = "Buzz"
		}

	}
	return text

}

func main() {

	for i := 1; i <= 100; i++ {
		getCond := FizzBuzz(i)

		fmt.Printf("%d -> %s \n", i, getCond)

	}

}
