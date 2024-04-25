package main

import "fmt"

func removeSliceElement(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func removeSwap(slice []int, i int) []int { // Если не важен порядок элементов
	slice[i] = slice[len(slice)-1] // Меняем местами последний элемент и элемент, который нам надо удалить
	return slice[:len(slice)-1]    // Возвращаем срез всех значений до последнего
}

func removeShift(slice []int, i int) []int { // Если порядок важен, но менее эффективно
	copy(slice[i:], slice[i+1:]) // Копируем все элементы после того, который нам надо удалить на его место
	return slice[:len(slice)-1]  // Так как размер не меняется, то последний элемент повторится 2 раза, так что срезваем его
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("removeSliceElement: ", removeSliceElement(slice1, 2))
	fmt.Println("removeSwap: ", removeSwap(slice2, 2))
	fmt.Println("removeShift: ", removeShift(slice3, 2))
}
