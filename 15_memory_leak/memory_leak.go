package main

import "strings"

func createHugeString(size int) string {
	str := make([]rune, size)
	return string(str)
}

var justString string

func someFunc() {
	// Функция выделяет 1мб памяти (1 на 10 позиции в двузначном числе 1024) на строку
	v := createHugeString(1 << 10)

	// Здесь мы хотим извлечь подстроку размером 100 рун,
	// но присваивая срез таким способом к justString мы продолжаем хранить указатель на оригинальную строку,
	// даже после завершения функции, тк justString- глобальная переменая и хранится в куче,
	// что приводит к тому, что сборщик мусора не освободит память выделенную на большую строку после завершения someFunc()
	//justString = v[:100]

	// Иначе- лучше скопировать подстроку
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
