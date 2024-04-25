package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int) // Обьявляем большие инты
	b := new(big.Int)

	a.SetString("1000000", 10) // Число из строки в число с основанием 10
	b.SetString("2000000", 10)

	sum := new(big.Int).Add(a, b) // Сложение
	fmt.Println(sum)

	sub := new(big.Int).Sub(a, b) // Вычитание
	fmt.Println(sub)

	mul := new(big.Int).Mul(a, b) // Умножение
	fmt.Println(mul)

	if b.Cmp(big.NewInt(0)) != 0 { // Если b == 0, то в ошибку
		div := new(big.Int).Div(a, b) // Делим, в данном случае ответ будет 0, тк отбросили дробную часть (инт)
		fmt.Println(div)
	} else {
		fmt.Println("Derived value is zero")
	}
}
