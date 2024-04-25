package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int, count int) { // Отправитель отправляет каждую секунду значение от 0 до N, передаем отправляющий канал и значение
	defer close(ch) // Закрыть канал перед выходом из функции
	for i := 0; i < count; i++ {
		ch <- i
		fmt.Println("Sent:", i)
		time.Sleep(time.Second)
	}
}

func receiver(ch <-chan int) { // Приемник принимает значение и печатает его в stdout, если что-то не так, то выходит из функции, передаем отправялющий канал
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("Channel closed. Exiting receiver.")
			return
		}
		fmt.Println("Received:", value)
	}
}

func main() {
	var seconds int
	fmt.Print("Введите количество секунд: ")
	fmt.Scanln(&seconds)

	ch := make(chan int) // Канал связи между отправителем и приемником

	go sender(ch, seconds) // Отправитель в гоурутину

	go receiver(ch) // Приемник в гоурутину

	time.Sleep(time.Duration(seconds) * time.Second) // Программа работает переданное количество секунд

	fmt.Println("Main goroutine exiting.")
}
