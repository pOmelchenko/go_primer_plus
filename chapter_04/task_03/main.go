package main

import "fmt"

// Упражнение по программированию 4.3
//
// Напишите программу, которая читает число с плавающей запятой и выводит его сначала в десятичной, а затем в
// экспоненциальной форме. Предусмотрите вывод в следующих форматах (количество цифр показателя степени в вашей системе
// может быть другим.
//
// Вводом является 21.3 или 2.1e+001
// Вводом является +21.290 или 2.129E+001
func main() {
	var number float32
	fmt.Print("Введите число с плавающей точкой: ")
	fmt.Scan(&number)

	fmt.Printf("Вводом является %.1f или %.1e\n", number, number)
	fmt.Printf("Вводом является %+.3f или %.3E\n", number, number)
}
