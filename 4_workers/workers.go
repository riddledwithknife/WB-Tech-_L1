package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	dataChan := make(chan string)
	var numWorkers int
	fmt.Println("Enter number of workers: ")
	fmt.Scanln(&numWorkers) // Получаем количество воркеров

	ctx, cancel := context.WithCancel(context.Background()) // Контекст прерывания работы

	sigChan := make(chan os.Signal, 1)   // Канал прерывания
	signal.Notify(sigChan, os.Interrupt) // Прерывание на ctrl+c

	go func() { // Гоурутина ждущая прерывание
		<-sigChan
		fmt.Println("Received interrupt, shutting down...")
		cancel() // Вызываем отмену контекста
	}()

	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerNum int) { // Создаем воркеров, которые принимают данные из мэйновой гоурутины и печатают данные. Завершают работу, если пришло прерывание
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("Worker %d exiting\n", workerNum)
					return
				case data := <-dataChan:
					fmt.Printf("Worker %d received data: %s\n", workerNum, data)
				}
			}
		}(i + 1)
	}

	for i := 0; ; i++ { // Цикл значений
		select {
		case <-ctx.Done():
			fmt.Printf("Main goroutine exiting\n")
			wg.Wait() // Ожидание завершения работы воркеров
			return
		default:
			data := fmt.Sprintf("Data %d", i)
			dataChan <- data // Отправка данных воркерам
		}
	}
}
