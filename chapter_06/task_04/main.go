package main

import "fmt"

const START_LETTER = 'A'

const ROWS = 6

// Упражнение по программированию 6.4
//
// Воспользуйтесь вложенными циклами, чтобы написать программу, которая выводит следующую последовательность
// символов:
// <pre>
// A<br>
// BC<br>
// DEF<br>
// GHIJ<br>
// KLMNO<br>
// PQRSTU
// </pre>
//
// @note Если в вашей системе не используется кодировка, в которой буквы представлены в числовом порядке, см. примечание
// в упражнении 3.
func main() {
	var (
		row    int32
		col    int32
		offset int32
	)

	offset = 0

	for row = 0; row < ROWS; row++ {
		for col = 0; col <= row; col++ {
			fmt.Printf("%c", START_LETTER+offset)
			offset++
		}

		fmt.Println()
	}
}
