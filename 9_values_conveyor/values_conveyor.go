package main

import "fmt"

func main() {
	numbers := make(chan int) // Создаем каналы
	results := make(chan int)

	go func() {
		arr := []int{1, 2, 3, 4, 5}
		for _, num := range arr { // Идем по массиву, отправляем значения в канал numbers
			numbers <- num
		}
		close(numbers) // Закрываем канал на отправку
	}()

	go func() {
		for num := range numbers { // Идем по каналу numbers
			results <- num * 2 // Отправляем в канал results значения, умноженные на 2
		}
		close(results) // Закрываем канал
	}()

	for result := range results { // Печатаем значения из канала
		fmt.Println(result)
	}
}
