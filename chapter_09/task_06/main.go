package main

import "fmt"

// Упражнение по программированию 9.6
//
// Напишите и протестируйте функцию, которая принимает в качестве аргументов адреса трех переменных
// <code>double</code> и помещает наименьшее значение в первую переменную, среднее значение — во вторую, а наибольшее
// значение — в третью.
func main() {
	var a, b, c float32

	orderNumbers(&a, &b, &c)

	fmt.Printf("a: %f, b: %f, c: %f", a, b, c)
}

func orderNumbers(a, b, c *float32) {
	minNum := *a

	if *b < minNum {
		minNum = *b
	}

	if *c < minNum {
		minNum = *c
	}

	awgNum := (*a + *b + *c) / 3

	maxNum := *a

	if *b > maxNum {
		maxNum = *b
	}

	if *c > maxNum {
		maxNum = *c
	}

	*a = minNum
	*b = awgNum
	*c = maxNum
}
