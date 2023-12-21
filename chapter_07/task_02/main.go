package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Упражнение по программированию 7.2
//
// @details Напишите программу, которая читает входные данные до тех пор, пока не встретится символ <code>#</code>.
// Программа должна выводить каждый введенный символ и его десятичный код <code>ASCII</code>. Каждая строка вывода
// должна содержать восемь пар "символ-код".
//
// @note Используйте счетчики символов и операцию деления по модулю (<code>%</code>) для вывода символа новой строки
// на каждой восьмой итерации цикла.
func main() {
	var (
		symbol, beakline rune
		counter          int
		err              error
	)

	reader := bufio.NewReader(os.Stdin)

	symbol, _, err = reader.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		return
	}
	for symbol != '#' {
		counter++

		if symbol == '\n' {
			continue
		}

		if counter%8 != 0 {
			beakline = ' '
		} else {
			beakline = '\n'
		}

		fmt.Printf("%c-%d%c", symbol, symbol, beakline)

		symbol, _, err = reader.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			return
		}
	}
}
