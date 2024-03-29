package main

import "fmt"

// Упражнение по программированию 9.3
//
// Напишите функцию, которая принимает три аргумента: символ и два целых числа. Символ предназначен для вывода.
// Первое целое значение задает количество указанных символов в строке, а второе целое число устанавливает количество
// таких строк. Напишите программу, в которой используется эта функция.
func main() {
	printChar('a', 10, 5)
}

func printChar(ch rune, cols, rows int) {
	var (
		c int
		r int
	)

	for r = 0; r < rows; r++ {
		for c = 0; c < cols; c++ {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}
