package main

import "fmt"

// Упражнение по программированию 7.9
//
// Напишите программу, которая принимает в качестве ввода положительное целое число и отображает все простые
// числа, которые меньше или равны введенному числу.
func main() {
	var (
		num   uint32
		div   uint32
		prime uint32
	)

	fmt.Print("Введите число: ")

	for {
		if _, err := fmt.Scan(&num); err != nil {
			fmt.Println("До свидания")
			return
		}

		count := 0
		for prime = 2; prime <= num; prime++ {
			isPrime := true

			for div = 2; div*div <= prime; div++ {
				if prime%div == 0 {
					isPrime = false
				}
			}

			if isPrime {
				count++
				fmt.Printf("%5d", prime)

				if count%10 == 0 {
					fmt.Println()
				}
			}
		}

		fmt.Println()
		fmt.Println("Введите следующее число для анализа, или введите q для завершения.")
	}
}
