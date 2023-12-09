package main

import "fmt"

// Упражнение по программированию 6.11
//
// Напишите программу, которая читает восемь целых чисел в массив, после чего выводит их в обратном порядке
func main() {
	var numbers [8]int

	fmt.Print("Введите 8 целых чисел: ")
	_, err := fmt.Scanf("%d%d%d%d%d%d%d%d",
		&numbers[0],
		&numbers[1],
		&numbers[2],
		&numbers[3],
		&numbers[4],
		&numbers[5],
		&numbers[6],
		&numbers[7])

	if err != nil {
		fmt.Println("Ввели какуюто дичь, а не число")
	}

	fmt.Printf(
		"%d%d%d%d%d%d%d%d",
		numbers[7],
		numbers[6],
		numbers[5],
		numbers[4],
		numbers[3],
		numbers[2],
		numbers[1],
		numbers[0])
}
