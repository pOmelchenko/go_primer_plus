package main

import "fmt"

// Упражнение по программированию 9.2
//
// Напишите функцию по имени <code>chline(ch, i, j)</code>, которая выводит требуемый символ в столбцах с
// <code>i</code> по <code>j</code>. Протестируйте эту функцию с помощью простого драйвера.
func main() {
	chline('a', 3, 3)
	chline('b', 3, 6)
	chline('c', 6, 3)
}

func chline(ch rune, i int, j int) {
	for y := 0; y < j; y++ {
		for x := 0; x < i; x++ {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}
