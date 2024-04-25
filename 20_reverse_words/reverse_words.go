package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s) // Делим на слова, аналог Split по пробелам
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем местами в ручную
	}

	//slices.Reverse(words) // Через метод

	return strings.Join(words, " ") // Обьединяем в строку с сепаратором- пробелом
}

func main() {
	words := "snow dog sun"

	fmt.Println(reverseWords(words))
}
