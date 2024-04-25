package main

import "fmt"

// Human структура
type Human struct {
	Name string
	Age  int
}

// Speak метод структуры Human
func (h *Human) Speak() {
	fmt.Printf("My name is %s, my age is %d\n", h.Name, h.Age)
}

// Action дочерняя структура структуры Human
type Action struct {
	Human //Встраиваем структуру Human и ее методы в Action
	Hobby string
}

// SomeAction метод структуры Action, которая в следствие встривания позволяет вызывать методы структуры Human
func (a *Action) SomeAction() {
	a.Speak()
	fmt.Printf("%s's hobby is %s", a.Name, a.Hobby)
}

func main() {
	a := Action{
		Human: Human{Name: "Daniil", Age: 21},
		Hobby: "skiing",
	}

	a.SomeAction()
}
