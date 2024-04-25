package main

import (
	"fmt"
	"reflect"
)

func getType(variable interface{}) string { // Принимаем интерфейс (любой тип данных)
	switch reflect.TypeOf(variable).Kind() { // Используем пакет reflect для определения типа данных
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "chan"
	default:
		return "Unhandled variable type"
	}
}

func main() {
	intVar := 10
	strVar := "hello"
	boolVar := true
	ch := make(chan int)

	// Печатаем
	fmt.Println(getType(intVar))
	fmt.Println(getType(strVar))
	fmt.Println(getType(boolVar))
	fmt.Println(getType(ch))
}
