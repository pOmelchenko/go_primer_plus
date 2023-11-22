package main

import "fmt"

// @brief Упражнение по программированию 5.8
//
// @details Напишите программу, которая выводит результаты применения операции деления по модулю. Пользователь должен
// первым ввести целочисленное значение, которое используется в качестве второго операнда и остается неизменным.
// Затем пользователь должен вводить числа, для которых будет вычисляться результат деления по модулю. Процесс должен
// прерываться вводом значения, которое равно или меньше 0. Пример выполнения этой программы должен выглядеть следующим
// образом:
//
// <pre>Эта программа вычисляет результаты деления по модулю.
// Введите целое число, которое будет служить вторым операндом: 256
// Теперь введите первый операнд: 438
// 438 % 256 равно 182
//
// Введите следующее число для первого операнда (<= 0 для выхода из программы): 1234567
// 1234567 % 256 равно 135
//
// Введите следующее число для первого операнда (<= 0 для выхода из программы): 0
// Готово</pre>
func main() {
	var (
		divider  int
		dividend int
	)

	fmt.Println("Эта программа вычисляет результаты деления по модулю.")
	fmt.Print("Введите целое число, которое будет служить вторым операндом: ")
	fmt.Scan(&divider)

	fmt.Print("Теперь введите первый операнд: ")
	fmt.Scan(&dividend)

	for dividend > 0 {
		fmt.Printf("%d %% %d результат: %d\n\n", dividend, divider, dividend%divider)
		fmt.Print("Введите следующее число для первого операнда (<= 0 для выхода из программы): ")
		fmt.Scan(&dividend)
	}

	fmt.Println("Готово")
}