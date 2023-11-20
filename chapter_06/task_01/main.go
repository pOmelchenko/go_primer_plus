package main

import "fmt"

const LETTERS_COUNT = 26

const START_LETTER = 'a'

// Упражнение по программированию 6.1
//
// Напишите программу, которая создает массив из 26 элементов и помещает в него 26 строчных букв английского
// алфавита. Также предусмотрите вывод содержимого этого массива.
func main() {
	var (
		chars [LETTERS_COUNT]int32
		i     int32
	)

	for i = 0; i < LETTERS_COUNT; i++ {
		chars[i] = START_LETTER + i
	}

	for i = 0; i < LETTERS_COUNT; i++ {
		fmt.Printf("%c", chars[i])
	}

	fmt.Println()
}
