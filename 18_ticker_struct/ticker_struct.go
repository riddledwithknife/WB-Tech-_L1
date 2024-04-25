package main

import (
	"fmt"
	"sync"
)

type Counter struct { // Структура счетчика гоурутин
	mu    sync.Mutex // Мьютекс для синхронизации доступа к памяти из гоурутин
	value int
}

func (c *Counter) Increment() { // Инкрементировать значение счетчика
	c.mu.Lock() // Блокируем мьютекс, чтобы не было race condition
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int { // Получить значение счетчика
	c.mu.Lock() // Блокируем
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := Counter{value: 0} // Инициируем структуру счетчика

	var wg sync.WaitGroup // Вэйт группа, чтобы отслеживать количество гоурутин

	numRoutines := 100 // Число гоурутин

	wg.Add(numRoutines) // Добавляем количество гоурутин в вэйт группу

	for i := 0; i < numRoutines; i++ {
		go func() { // Инкрементируем счетчик в каждой гоурутине
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Ожидаем выполнение всех гоурутин

	fmt.Println("Значение счетчика: ", counter.Value()) // Печатаем
}
