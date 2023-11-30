package main

import "fmt"

// Упражнение по программированию 6.7
//
// Напишите программу, которая читает слово в символьный массив, а затем выводит это слово в обратном порядке.
//
// Воспользуйтесь функцией <code>strlen()</code> (глава 4) для вычисления индекса последнего символа массива.
func main() {
	var (
		word  string
		index int
	)

	fmt.Print("Введите слово: ")
	_, err := fmt.Scan(&word)

	if err != nil {
		fmt.Println("Фигня какая-то, а не строка")
	}

	for index = len(word) - 1; index >= 0; index-- {
		fmt.Printf("%c", word[index])
	}
}
