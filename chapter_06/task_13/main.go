package main

import "fmt"

const ELEMENTS_COUNT int = 8

// Упражнение по программированию 6.13
//
// Напишите программу, которая создает восьмиэлементный массив типа <code>int</code> и помещает в него элементы
// начальных восьми степеней <code>2</code>, а затем выводит полученные значения. Применяйте цикл <code>for</code> для
// вычисления элементов массива, и ради разнообразия для отображения значений воспользуйтесь циклом
// <code>do while</code>.
func main() {
	var (
		results       [ELEMENTS_COUNT]int
		element_index int
		power         int
	)

	for element_index = 0; element_index < ELEMENTS_COUNT; element_index++ {
		results[element_index] = 2

		for power = 2; power <= element_index+1; power++ {
			results[element_index] = results[element_index] * 2
		}
	}

	element_index = 0

	for element_index < ELEMENTS_COUNT {
		fmt.Printf("%d\n", results[element_index])
		element_index++
	}
}
