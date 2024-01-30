package main

import "fmt"

// Упражнение по программированию 9.5
//
// Напишите и протестируйте функцию по имени <code>large_of()</code>, которая заменяет содержимое двух
// переменных <code>double</code> большим из значений. Например, вызов <code>large_of(x,y)</code> присвоит переменным
// <code>x</code> и <code>y</code> большее из их значений.
func main() {
	var a, b float32

	a = 10.1
	b = 3.2

	largeOf(&a, &b)

	fmt.Printf("Result a = %0.2f, b = %0.2f", a, b)
}

func largeOf(a *float32, b *float32) {
	if *a > *b {
		*b = *a
	} else {
		*a = *b
	}
}
