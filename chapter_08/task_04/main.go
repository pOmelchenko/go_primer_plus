package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Упражнение по программированию 8.4
//
// Напишите программу, которая читает ввод как поток символов, пока не встретит <code>EOF</code>. Программа
// должна сообщать среднее количество букв в словах. Не считайте пробельные символы в словах буквами. На самом деле, так
// же не должны учитываться и знаки препинания, но в данном упражнении об этом можно не беспокоиться. (Для учета знаков
// препинания можно воспользоваться функцией <code>ispunct()</code> из семейства <code>ctype.h</code>.)
func main() {
	var (
		words        int
		charsInWords int
		word         bool
	)

	scanner := bufio.NewReader(os.Stdin)

	for {
		input, err := scanner.ReadByte()
		if err == io.EOF && word == false {
			break
		}

		if word {
			if unicode.IsLetter(rune(input)) {
				word = true
				charsInWords++
			} else {
				word = false
				words++
			}
		} else {
			if unicode.IsLetter(rune(input)) {
				word = true
				charsInWords++
			}
		}

		if err == io.EOF && word == true {
			words++
			break
		}
	}

	fmt.Printf(
		"Слов: %d\nВсего символов: %d\nСреднне количество символов в словах: %d\n",
		words,
		charsInWords,
		charsInWords/words)
}
