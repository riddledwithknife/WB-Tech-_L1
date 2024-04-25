package main

import (
	"fmt"
	"math"
)

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groupedTemperatures := make(map[int][]float64)

	for _, temp := range temperatures { // Идем по температурам
		groupKey := 0

		if temp >= 0 { // Вычиляем группу температур
			groupKey = int(math.Floor(temp/10)) * 10
		} else {
			groupKey = int(math.Floor(temp/10)+1) * 10 // Если меньше нуля, то прибавляем единицу из-за округления
		}

		groupedTemperatures[groupKey] = append(groupedTemperatures[groupKey], temp)
	}

	return groupedTemperatures
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // Массив температур

	groupedTemperatures := groupTemperatures(temperatures)

	fmt.Println(groupedTemperatures) // Печатаем в stdout
}
