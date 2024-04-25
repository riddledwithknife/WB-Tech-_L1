package main

import (
	"fmt"
	"sync"
)

type ConcDict struct { // Структура конкуррентной мапы
	mu   sync.Mutex
	data map[string]int
}

func NewConcDict() *ConcDict { // Конструктор
	return &ConcDict{
		data: make(map[string]int),
	}
}

func (c *ConcDict) Set(key string, value int) { // Запись значения
	c.mu.Lock() // Блокируем мьютекс, чтобы не получить race condition
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *ConcDict) Get(key string) (int, bool) { // Взятие значения
	c.mu.Lock() // Блокируем мьютекс
	defer c.mu.Unlock()
	value, ok := c.data[key]
	return value, ok
}

func main() {
	dict := NewConcDict() // Создаем мапу

	var wg sync.WaitGroup // Вэйт группа
	for i := 0; i < 10; i++ {
		wg.Add(1) // Добавляем гоурутину в сетчик
		go func(i int) {
			defer wg.Done() // Откладываем выполнения учета выполненной гоурутины
			key := fmt.Sprintf("key-%d", i)
			dict.Set(key, i)
			value, _ := dict.Get(key)
			fmt.Printf("key: %s, value: %d\n", key, value)
		}(i)
	}
	wg.Wait()
}
