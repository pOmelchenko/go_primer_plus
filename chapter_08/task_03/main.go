package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// Упражнение по программированию 8.3
//
// Напишите программу, которая читает ввод как поток символов, пока не встретит <code>EOF</code>. Программа
// должна сообщать количество прописных букв, количество строчных букв и количество остальных символов во входных
// данных. Можете предполагать, что числовые значения для строчных букв являются последовательными, и то же самое
// справедливо для прописных букв. Либо для большей переносимости можете использовать подходящие классификационные
// функции из библиотеки <code>ctype.h</code>.
func main() {
	var (
		lower int
		upper int
		etc   int
	)

	scanner := bufio.NewReader(os.Stdin)

	for {
		input, err := scanner.ReadByte()

		if err == io.EOF {
			break
		}

		if unicode.IsLower(rune(input)) {
			upper++
		} else if unicode.IsUpper(rune(input)) {
			lower++
		} else {
			etc++
		}
	}

	fmt.Printf("Верхний регистр: %d\nНижний регистр: %d\nОстальные: %d\n", upper, lower, etc)
}
