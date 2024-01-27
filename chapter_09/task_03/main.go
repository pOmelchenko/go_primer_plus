package main

import "fmt"

func main() {
	printChar('a', 10, 5)
}

func printChar(ch rune, cols, rows int) {
	var (
		c int
		r int
	)

	for r = 0; r < rows; r++ {
		for c = 0; c < cols; c++ {
			fmt.Print(string(ch))
		}
		fmt.Println()
	}
}
