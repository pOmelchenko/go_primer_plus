package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Упражнение по программированию 8.1
//
// Напишите программу, которая подсчитывает количество символов во входных данных до достижения конца файла.
func main() {
	var count int

	scanner := bufio.NewReader(os.Stdin)

	for {
		_, err := scanner.ReadByte()
		if err == io.EOF {
			break
		}
		count++
	}

	fmt.Printf("В файле было %d символов.\n", count)
}
