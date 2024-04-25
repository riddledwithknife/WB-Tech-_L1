package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []int{5, 4, 3, 6, 8, 10, 20, 17, 6}
	slices.Sort(slice) // Сортиурем слайс
	fmt.Println(slice) // Печатаем
}
