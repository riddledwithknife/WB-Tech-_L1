package main

import (
	"fmt"
	"math"
)

type Point struct { // Структура точки
	x, y float64
}

func NewPoint(x, y float64) Point { // Конструктор
	return Point{x, y}
}

func (p Point) Distance(q Point) float64 {
	return math.Sqrt(math.Pow(q.x-p.x, 2) + math.Pow(q.y-p.y, 2))
}

func main() {
	point1 := NewPoint(-1, 3)
	point2 := NewPoint(6, 2)

	distance := point1.Distance(point2)

	fmt.Println(distance)
}
