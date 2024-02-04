package main

import (
	"fmt"
	"os"
)

// Упражнение по программированию 9.10
//
// Обобщите функцию <code>to_binary()</code> из листинга 9.8 до функции <code>to_base_n()</code>, которая
// принимает второй аргумент в диапазоне от 2 о 10. Она должна выводить число, переданное в первом аргументе, в системе
// счисления с основанием, которое указанно во втором аргументе. Например, вызов <code>to_base_n(129,8)</code> должен
// отобразить <code>201</code>, т.е. восьмеричный эквивалент числа <code>129</code>. Протестируйте готовую функцию в
// какой-нибудь программе.
func main() {
	var (
		number int
		base   int
		err    error
	)

	fmt.Println("Введите целое число (q для завершения):")

	_, err = fmt.Scanf("%d%d", &number, &base)

	if err != nil {
		fmt.Println("Программа завершена.")
		os.Exit(0)
	}

	for {
		fmt.Println("Двоичный эквивалент: ")
		toBinary(number, base)
		fmt.Println()
		fmt.Println("Введите целое число (q для завершения):")

		_, err = fmt.Scanf("%d%d", &number, &base)

		if err != nil {
			break
		}
	}
	fmt.Println("Программа завершена.")
}

func toBinary(number, base int) {
	if base < 2 || base > 10 {
		fmt.Println("!!! Задан неверный диапазон основания системы счисления")
		return
	}

	var result int

	result = number % base

	if number >= base {
		toBinary(number/base, base)
	}

	fmt.Printf("%d", result)
}
