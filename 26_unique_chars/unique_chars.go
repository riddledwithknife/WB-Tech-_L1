package main

import (
	"fmt"
	"strings"
)

func areAllCharsUnique(str string) bool {
	str = strings.ToLower(str) // Переводим все руны в нижний регистр

	charMap := make(map[rune]bool) // Мапа рун

	for _, char := range str { // Идем по строке
		if _, ok := charMap[char]; ok { // Если такой символ уже есть, то возвращаем false
			return false
		}
		charMap[char] = true // Записываем символ в мапу
	}

	return true
}

func main() {
	fmt.Println(areAllCharsUnique("abcd"))
	fmt.Println(areAllCharsUnique("abCdefAaf"))
	fmt.Println(areAllCharsUnique("aabcd"))
}
