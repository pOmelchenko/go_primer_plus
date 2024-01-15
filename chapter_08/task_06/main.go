package main

import (
	"bufio"
	"fmt"
	"os"
)

// Упражнение по программированию 8.6
//
// Модифицируйте функцию <code>get_first()</code> из листинга 8.8 так, чтобы она возвращала первый встречный
// непробельный символ. Протестируйте ее в какой-нибудь простой программе.
func main() {
	var ch rune
	ch = getFirst()

	fmt.Printf("Результат: %c\n", ch)
	fmt.Println("Программа завершена.")
}

func getFirst() rune {
	scanner := bufio.NewReader(os.Stdin)

	for {
		choice, _, _ := scanner.ReadRune()

		if choice == ' ' || choice == '\n' {
			continue
		}

		return choice
	}
}
