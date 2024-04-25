package main

import "fmt"

func intersection(set1, set2 []int) []int {
	set1Map := make(map[int]bool) // Мапа 1ого множества

	for _, val := range set1 { // Заоплняем значениями
		set1Map[val] = true
	}

	intersectionMap := make(map[int]bool) // Мапа пересечения
	for _, val := range set2 {
		if _, ok := set1Map[val]; ok { // Если в 1 множестве есть такое значение, которое есть во втором, то заполняем мапу пересечения
			intersectionMap[val] = true
		}
	}

	intersectionArr := make([]int, 0, len(intersectionMap)) // Из мапы в слайс
	for k := range intersectionMap {
		intersectionArr = append(intersectionArr, k)
	}
	return intersectionArr
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 6} // Два множества
	set2 := []int{2, 4, 6, 9, 10}

	result := intersection(set1, set2)
	fmt.Println(result) // Печатаем
}
