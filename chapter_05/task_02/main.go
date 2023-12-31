package main

import "fmt"

// Упражнение по программированию 5.2
//
// Напишите программу, которая запрашивает у пользователя ввод целого числа, а затем выводит все целые числа,
// начиная с этого числа (включая его) и заканчивая числом, которое больше введенного значения на 10 (включая его).
// То есть,если вводится число 5, то в выводе должны присутствовать числа от 5 до 15.) Обеспечьте разделение выводимых
// значений друг от друга пробелами, символами табуляции или символами новой строки.
func main() {
	var start int
	var finish int
	fmt.Printf("Введите число: ")
	fmt.Scan(&start)

	finish = start + 10

	for start <= finish {
		fmt.Println(start)
		start++
	}
}
