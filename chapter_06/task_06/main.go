package main

import "fmt"

// Упражнение по программированию 6.6
//
// Напишите программу для вывода таблицы, в каждой строке которой представлено целое число, его и его куб.
// Запросите у пользователя верхний и нижний пределы таблицы. Используйте цикл <code>for</code>.
func main() {
	var (
		minNum int
		maxNum int
		row    int
	)

	fmt.Print("Введите нижний и верхний пределы таблицы: ")
	fmt.Scan(&minNum, &maxNum)

	for row = minNum; row <= maxNum; row++ {
		fmt.Printf(
			"%5d | %5d | %5d\n",
			row,
			row*row,
			row*row*row)
	}
}
