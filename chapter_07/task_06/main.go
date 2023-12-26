package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Упражнение по программированию 7.6
//
// Напишите программу, которая читает входные данные, пока не встретит символ <code>#</code>,
// и сообщает количество вхождений последовательности <code>ei</code>.
//
// Программа должна "запоминать" предыдущий символ, а также текущий символ. Проверьте ее на входной
// последовательности вроде "Receive your eieio award#".
func main() {
	var (
		last     rune
		prew     rune
		err      error
		ei_count int
	)

	reader := bufio.NewReader(os.Stdin)

	last, _, err = reader.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		return
	}

	for last != '#' {
		if 'e' == prew && 'i' == last {
			ei_count++
		}

		prew = last

		last, _, err = reader.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			return
		}
	}

	fmt.Printf("'ei' встретилось %d раз", ei_count)
}
