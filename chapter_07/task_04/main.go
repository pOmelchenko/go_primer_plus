package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Упражнение по программированию 7.4
//
// Используя операторы <code>if else</code>, напишите программу, которая читает входные данные, пока не
// встретит символ <code>#</code>, заменяет каждую точку восклицательным знаком, изначально присутствующие
// восклицательные знаки — двумя восклицательными знаками и в конце сообщает о количестве произведенных замен.
func main() {
	var (
		input   rune
		counter int
		err     error
	)

	reader := bufio.NewReader(os.Stdin)

	input, _, err = reader.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		return
	}

	for input != '#' {
		if input == '.' || input == '!' {
			if input == '!' {
				fmt.Printf("%s", "!!")
			} else {
				fmt.Printf("%s", "!")
			}

			counter++
		} else {
			fmt.Printf("%s", string(input))
		}

		input, _, err = reader.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			return
		}
	}

	fmt.Printf("\nвсего замен: %d\n", counter)
}
