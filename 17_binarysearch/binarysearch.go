package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	pos, in := slices.BinarySearch(slice, 5) // Бинарный поиск
	fmt.Println(slice, "\n", pos, "\n", in)  // Печатаем слайс, позицию таргета, и есть ли вообще этот элемент в слайсе
}
