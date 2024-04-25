package main

import (
	"fmt"
	"sync"
)

func square(wg *sync.WaitGroup, val int, ch chan int) {
	defer wg.Done() // Отложенный вызов, сигнализирующий о заверешении гоурутины
	squared := val * val
	ch <- squared // Отправляем значения квадрата в канал
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup          // Вэйт группа
	ch := make(chan int, len(arr)) // Интовый канал размером длины массива

	for _, value := range arr {
		wg.Add(1)                 // Добавляем гоурутины в вэйтгруппу
		go square(&wg, value, ch) // Вызов возведения в квадрат в гоурутине
	}

	wg.Wait() // Ожидаем завершения всех гоурутин
	close(ch) // Закрываем канал

	for result := range ch {
		fmt.Println(result)
	}
}
