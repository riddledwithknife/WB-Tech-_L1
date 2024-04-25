package main

import "fmt"

type LegacyPrinter interface { // Устаревший интерфейс
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (p *MyLegacyPrinter) Print(s string) string {
	return fmt.Sprintf("From legacy printer: %s", s)
}

type NewPrinter interface { // Новый интерфейс
	NewPrint(s string) string
}

type MyNewPrinter struct { // Встраиваем старый интерфейс в новую структуру
	LegacyPrinter
}

func (p *MyNewPrinter) NewPrint(s string) string { // Имеем возможность вызвать старый код из нового без изменения старого
	return p.Print("Adapter is printing")
}

func main() {
	legacyPrinter := &MyLegacyPrinter{}

	adapter := &MyNewPrinter{
		LegacyPrinter: legacyPrinter,
	}

	result := adapter.NewPrint("From legacy printer")
	fmt.Println(result)
}
