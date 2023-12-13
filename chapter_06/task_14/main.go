package main

import "fmt"

const ARRAY_LIMIT = 8

// Упражнение по программированию 6.14
//
// Напишите программу, которая создает два восьмиэлементных массива типа <code>double</code> и использует цикл
// для ввода значений восьми элементов первого массива. Программа должна накапливать в элементах второго массива суммы
// первого массива с нарастающим итогом. Например, четвертый элемент второго массива должен быть равен сумме первых
// четырех элементов первого массива, а пятый элемент второго второго массива — сумме пяти первых элементов первого
// массива. (Это можно сделать с помощью вложенных циклов, однако если учесть тот факт, что пятый элемент второго
// массива равен четвертому элементу второго массива плюс пятый элемент первого массива, можно избежать вложенных циклов
// и применить для решения задачи единственный цикл.) В завершение воспользуйтесь циклом для вывода содержимого обоих
// массивов, при этом первый массив должен отображаться в первой строке, а каждый элемент второго массива должен
// помещаться непосредственно под соответствующим элементом первого массива.
func main() {
	var (
		array_input   [ARRAY_LIMIT]float32
		array_sum     [ARRAY_LIMIT]float32
		element_index int
	)

	fmt.Printf("Ведите значение для [%d] элемента массива: ", element_index)
	_, err := fmt.Scan(&array_input[element_index])

	if err != nil {
		fmt.Println("Введено что-то не правильное")
	}

	array_sum[element_index] = array_input[element_index]

	element_index++

	for element_index < ARRAY_LIMIT {
		fmt.Printf("Ведите значение для [%d] элемента массива: ", element_index)
		_, err := fmt.Scan(&array_input[element_index])

		if err != nil {
			fmt.Println("Введено что-то не правильное")
		}

		array_sum[element_index] = array_input[element_index] + array_sum[element_index-1]

		element_index++
	}

	for element_index = 0; element_index < ARRAY_LIMIT; element_index++ {
		fmt.Printf("%12f", array_input[element_index])
	}
	fmt.Println()

	for element_index = 0; element_index < ARRAY_LIMIT; element_index++ {
		fmt.Printf("%12f", array_sum[element_index])
	}
	fmt.Println()
}
