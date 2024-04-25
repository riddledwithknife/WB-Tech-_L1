package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second) // Функция возвращает канал, так что принимаем его в функции
}

func main() {
	fmt.Println("5 seconds sleep...")
	Sleep(5) // Вызов
}
