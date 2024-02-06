package main

import "fmt"

// Упражнение по программированию 9.11
//
// Напишите и протестируйте функцию <code>Fibonacci()</code>, в которой для вычисления чисел Фибоначчи вместо
// рекурсии применяется цикл.
func main() {
	var number = 7

	fmt.Printf("%d число в последовательности Фибоначчи: %d\n", number, Fibonacci(number, true))
}

func Fibonacci(number int, showSq bool) int {
	if number == 0 {
		return 0
	}

	if number == 1 || number == 2 {
		return 1
	}

	if showSq {
		fmt.Printf("1 1 ")
	}

	var (
		lastNum = 1
		result  = 1
		tmp     int
	)

	for step := 3; step <= number; step++ {
		tmp = result + lastNum
		lastNum = result
		result = tmp
		if showSq {
			fmt.Printf("%d ", result)
		}
	}

	if showSq {
		fmt.Println()
	}

	return result
}
