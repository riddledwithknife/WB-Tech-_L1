package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	var L int         // Левый указатель
	R := len(arr) - 1 // Правый указатель

	for L <= R {
		mid := (L + R) / 2 // Центр

		if target > arr[mid] { // Если искомое значение больше значения по середине, то сдвигаем левый указатель
			L = mid + 1
		} else if target < arr[mid] { // Если искомое значение меньше значения по середине, то сдвигаем правый указатель
			R = mid - 1
		} else {
			return mid
		}
	}

	return -1 // Если такого значения нет, то возвращаем -1
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	pos := binarySearch(slice, 4)

	fmt.Println(pos)
}
