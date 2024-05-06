package main

import (
	"fmt"
)

func quickSort(arr []int, s, e int) []int { // Принимаем слайс, начальный индекс и количество элементов
	if e-s+1 <= 1 { // Если количство элементов равно 1 или нулю, то возвращаем слайс
		return arr
	}

	pivot := arr[e] // Опорный элемент (Последний)
	left := s       // Левый указатель

	for i := s; i < e; i++ {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i] // Обмен переменной
			left++                                // Сдвиг левого указателя
		}
	}

	arr[e] = arr[left]
	arr[left] = pivot

	quickSort(arr, s, left-1) // Рекурсивный вызов левой части
	quickSort(arr, left+1, e) // Рекурсивный вызов правой части

	return arr
}

func main() {
	slice := []int{5, 4, 3, 6, 8, 10, 20, 17, 6}

	slice = quickSort(slice, 0, len(slice)-1)

	fmt.Println(slice) // Печатаем
}
