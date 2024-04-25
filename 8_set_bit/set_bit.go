package main

import (
	"fmt"
)

func setBit(n int64, i int, value int) int64 {
	// Проверяем, что i находится в допустимом диапазоне
	if i < 0 || i > 63 { // 64й бит- знаковый
		fmt.Println("Wrong index")
		return n
	}

	// Создаем маску для установки бита в нужном месте
	mask := int64(1 << uint(i))

	if value == 1 {
		// Устанавливаем i-й бит в 1
		n |= mask // Поразрядное ИЛИ (Если в маске будет единица, а в числе- 0, то поменяется на 1)
	} else if value == 0 {
		// Устанавливаем i-й бит в 0
		n &= ^mask // Сброс бита (Инвертирует маску, а потом логическое И)
	} else {
		fmt.Println("Wrong value")
		return n
	}

	return n
}

func main() {
	num := int64(42) // Пример числа типа int64
	i := 1           // Место бита
	value := 0       // Заначение бита

	// Устанавливаем i-й бит в указанное значение
	result := setBit(num, i, value)

	fmt.Println("Result: ", result)
}
