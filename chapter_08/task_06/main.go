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
	ch = get_first()

	fmt.Printf("Результат: %c\n", ch)
	fmt.Println("Программа завершена.")
}

func get_first() rune {
	scanner := bufio.NewReader(os.Stdin)
	for {
		choice, _ := scanner.ReadByte()
		if rune(choice) == ' ' {
			continue
		}

		if rune(choice) != '\n' {
			continue
		}

		return rune(choice)
	}
}
