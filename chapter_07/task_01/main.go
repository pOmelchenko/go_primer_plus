package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const STOP_SHAR = '#'

// Упражнение по программированию 7.1
//
// Напишите программу, которая читает входные данные до тех пор, пока не встретится символ <code>#</code>, а
// затем отображается количество прочитанных пробелов, количество символов новой строки и количество всех остальных
// символов.
func main() {
	var (
		spaces_count       int
		new_line_count     int
		other_charts_count int
		character          rune
		err                error
	)

	reader := bufio.NewReader(os.Stdin)

	character, _, err = reader.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		return
	}
	for character != STOP_SHAR {
		if character == ' ' {
			spaces_count++
		} else if character == '\n' {
			new_line_count++
		} else {
			other_charts_count++
		}
		character, _, err = reader.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			return
		}
	}

	fmt.Printf("Пробелов: %d\nНовых строк: %d\nПрочие символы: %d\n", spaces_count, new_line_count, other_charts_count)
}
