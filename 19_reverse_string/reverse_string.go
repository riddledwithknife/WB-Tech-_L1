package main

import (
	"fmt"
)

func reverseString(s string) string {
	runes := []rune(s) // Конвертируем в массив рун
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Меняем местами руны из начала и конца массива в ручную
	}

	//slices.Reverse(runes) // Через метод

	return string(runes) // Возвращаем как строку
}

func main() {
	str := "главрыба"

	revStr := reverseString(str)

	fmt.Println(revStr)
}
