package main

import (
	"context"
	"fmt"
	"time"
)

func workerBool(done chan bool) {
	fmt.Println("Bool goroutine starting...")

	for {
		select {
		case <-done: // Ждем, когда придет true из основной гоурутины
			fmt.Println("Bool goroutine end...")
			return
		}
	}
}

func workerCtx(ctx context.Context) {
	fmt.Println("Context goroutine starting...")

	for {
		select {
		case <-ctx.Done(): // Ждем, когда придет контекст из основной гоурутины
			fmt.Println("Context goroutine end...")
			return
		}
	}
}

func main() {
	done := make(chan bool)

	go workerBool(done)

	time.Sleep(time.Second) // Даем время на "работу"

	done <- true // Отправлем true в канал
	close(done)  // Закрываем канал

	ctx, cancel := context.WithCancel(context.Background()) // Создаем контекст

	go workerCtx(ctx)

	time.Sleep(time.Second) // Даем время на "работу"

	cancel()                // Прерываем
	time.Sleep(time.Second) // Ожидаем завершения
}
