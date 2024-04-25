package main

import "fmt"

func main() {
	set := make(map[string]bool) // Dummy мапа c ключами- нужными строками

	strArr := []string{"cat", "cat", "dog", "cat", "tree"} // Последовательность строк

	for _, str := range strArr { // Заполняем
		set[str] = true
	}

	fmt.Println(set) // Печатаем сет
}
