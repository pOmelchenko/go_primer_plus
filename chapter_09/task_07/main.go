package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Упражнение по программированию 9.7
//
// Напишите программу, которая читает символы из стандартного ввода вплоть до конца файла. Для каждого символа
// программа должна сообщать, является ли он буквой. Если символ — буква, программа вдобавок должна сообщать ее
// порядковый номер в алфавите. Например, буквы <code>c</code> и <code>C</code> будут иметь номер 3. Предусмотрите в
// программе функцию, которая принимает символ в качестве аргумента и возвращает его порядковый номер в алфавите, если
// он является буквой, и <code>-1</code> в противном случае.
func main() {
	var ch rune
	var err error

	scanner := bufio.NewReader(os.Stdin)

	for {
		ch, _, err = scanner.ReadRune()
		if err == io.EOF {
			break
		}

		fmt.Printf("%c - %d\n", ch, getRuneNumber(ch))
		ch, _, err = scanner.ReadRune()
	}
}

func getRuneNumber(ch rune) int {
	if unicode.IsLetter(ch) {
		return int(unicode.ToLower(ch))
	} else {
		return -1
	}
}
