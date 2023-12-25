package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Упражнение по программированию 7.5
//
// Выполните упражнение 4, но с применением оператора <code>switch</code>.
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

		switch input {
		case '.':
			fallthrough
		case '!':
			if input == '!' {
				fmt.Printf("%s", "!!")
			} else {
				fmt.Printf("%s", "!")
			}
			counter++
		default:
			fmt.Printf("%c", input)
		}

		input, _, err = reader.ReadRune()
		if err != nil {
			log.Println("stdin:", err)
			return
		}
	}

	fmt.Printf("\nвсего замен: %d\n", counter)
}
