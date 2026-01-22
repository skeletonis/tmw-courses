package main

import "fmt"

func main(){

	isTrue := true
	isFalse := false
	var result bool

	result = true && false
	fmt.Println(isTrue, " && ", isFalse, " = ", result)

	result = true && true
	fmt.Println(isTrue, " && ", isTrue, " = ", result)

result = false && false
	fmt.Println(isFalse, " && ", isFalse, " = ", result)

result = true || false
	fmt.Println(isTrue, " || ", isFalse, " = ", result)

result = true || true
	fmt.Println(isTrue, " || ", isTrue, " = ", result)

result = false || false
	fmt.Println(isFalse, " || ", isFalse, " = ", result)

	
}
