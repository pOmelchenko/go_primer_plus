package main

import "fmt"

// Упражнение по программированию 5.7
//
// Напишите программу, которая запрашивает ввод числа типа <code>double</code> и выводит значение куба этого
// числа. Для этого используйте собственную функцию, которая возводит значение в куб и выводит полученный результат.
// Программа <code>main()</code> должна передавать этой функции вводимое значение.
func main() {
	var number float64
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	fmt.Printf("Результат %.2f\n", cube(number))
}

func cube(number float64) float64 {
	return number * number * number
}
