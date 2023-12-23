package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Упражнение по программированию 7.3
//
// Напишите программу, которая читает целые числа до тех пор, пока не встретится число 0. После прекращения
// ввода программа должна сообщить обще количество введенных четных чисел (за исключением 0), среднее значение
// введенных четных чисел, общее количество введенных нечетных чисел и среднее значение нечетных чисел.
func main() {
	var (
		symbol  rune
		counter int
		sum     int
		even    int
		odd     int
		err     error
	)

	reader := bufio.NewReader(os.Stdin)

	symbol, _, err = reader.ReadRune()
	if err != nil {
		log.Println("stdin:", err)
		return
	}

	for symbol != '0' {
		counter++
		num, _ := strconv.Atoi(string(symbol))
		sum += num

		if num%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	fmt.Printf("Введено %d чисел, среднее число %d, четных %d, нечетных %d", counter, sum/counter, even, odd)
}
