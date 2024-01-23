package main

import "fmt"

// Упражнение по программированию 9.1
//
// Напишите функцию по имени <code>minValue(x,y)</code>, которая возвращает меньшее из двух значений
// <code>double</code>. Протестируйте эту функцию с помощью простого драйвера.
func main() {
	var (
		a float32
		b float32
	)

	a = 2.2
	b = 1.2

	fmt.Printf("a = %0.2f, b = %0.2f, res = %0.2f\n", a, b, minValue(a, b))

	a = 2.2
	b = 3.2

	fmt.Printf("a = %0.2f, b = %0.2f, res = %0.2f\n", a, b, minValue(a, b))
}

func minValue(a float32, b float32) float32 {
	if a < b {
		return a
	} else {
		return b
	}
}
