package main

import "fmt"

func main() {
	a := 5
	b := 10

	a, b = b, a // Обмен

	fmt.Println("a after vice versa: ", a)
	fmt.Println("b after vice versa: ", b)
}
