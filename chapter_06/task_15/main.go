package main

import "fmt"

const LIMIT = 10

// Упражнение по программированию 6.15
//
// Напишите программу, которая читает строку ввода, а затем выводит ее в обратном порядке. Ввод можно сохранять
// в массива значений типа <code>char</code>; предполагается, что строка состоит не более чем из 255 символов.
// Вспомните, что для чтения символа за раз можно применять функцию <code>scanf()</code> со спецификатором
// <code>%c</code>, а при нажатии клавиши <code><Enter></code> генерируется символ новой строки (<code>\\n</code>).
func main() {
	var (
		chars [LIMIT]byte
		index int
	)

	_, err := fmt.Scanf("%c", &chars[0])
	if err != nil {
		fmt.Print("Что-то не то, не символ")
	}

	for index = 1; index < LIMIT; index++ {
		_, err := fmt.Scanf("%c", &chars[index])
		if err != nil {
			fmt.Print("Что-то не то, не символ")
			break
		}
	}

	for index = LIMIT - 1; index >= 0; index-- {
		fmt.Printf("%c", chars[index])
	}

	fmt.Println()
}
