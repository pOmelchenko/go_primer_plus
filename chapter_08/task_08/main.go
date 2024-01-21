package main

import (
	"bufio"
	"fmt"
	"os"
)

// Упражнение по программированию 8.8
//
// Напишите программу, которая выводит на экран меню, предлагающее выбрать сложение, вычитание, умножение или
// деление. После выбора программа должна запросить два числа и затем выполнить затребованную операцию. Программа должна
// принимать только варианты, предлагаемые в меню. Для чисел должен использоваться тип <code>float</code> и программа
// должна предоставлять пользователю возможность повторно вводить числа, если он введет нечисловые данные. В случае
// деления программа должна предложить пользователю ввести другое значение, если он ввел для второго операнда значение
// <code>0</code>. Выполнение такой программы должно иметь примерно такой вид:
// <pre>
// Выберите желаемую операцию:<br>
// a. сложение            s. вычитание<br>
// m. умножение           d. деление<br>
// q. завершение<br>
// <b>с</b><br>
// Введите первое число: <b>22.4</b><br>
// Введите второе число: <b>один</b><br>
// один не является числом.<br>
// Введите число, такое как 2.5, -1.78E8, или 3: <b>1</b><br>
// 22.4 + 1 = 23.4<br>
// a. сложение            s. вычитание<br>
// m. умножение           d. деление<br>
// q. завершение<br>
// <b>д</b><br>
// Введите первое число: <b>18.4</b><br>
// Введите второе число: <b>0</b><br>
// Введите число отличное от 0: <b>0.2</b><br>
// 18.4 / 0.2 = 92<br>
// a. сложение            s. вычитание<br>
// m. умножение           d. деление<br>
// q. завершение<br>
// <b>з</b><br>
// Программа завершена.<br>
// </pre>
func main() {
	choice := getChoice()

	for choice != 'q' {
		fmt.Print("Введите первое число: ")
		numA := getNumber()

		fmt.Print("Введите второе число: ")
		numB := getNumber()

		switch choice {
		case 'a':
			fmt.Printf("%0.1f + %0.1f = %0.1f\n", numA, numB, numA+numB)
		case 's':
			fmt.Printf("%0.1f - %0.1f = %0.1f\n", numA, numB, numA-numB)
		case 'm':
			fmt.Printf("%0.1f * %0.1f = %0.1f\n", numA, numB, numA*numB)
		case 'd':
			if 0 == numB {
				fmt.Print("Введите число отличное от 0: ")
				numB = getNumber()
			}
			fmt.Printf("%0.1f / %0.1f = %0.1f\n", numA, numB, numA/numB)
		default:
			choice = getChoice()
		}
	}

	fmt.Println("Программа завершена.")
}

func getChoice() rune {
	fmt.Println("Выберите операцию из списка:")
	fmt.Println("a. сложение            s. вычитание")
	fmt.Println("m. умножение           d. деление")
	fmt.Println("q. завершение")

	choice := getFirst()

	for {
		if choice != 'a' && choice != 's' && choice != 'm' && choice != 'd' && choice != 'q' {
			fmt.Println("Выбран неверный вариант, попробуйте еще раз")
			choice = getFirst()
		} else {
			return choice
		}
	}

}

func getFirst() rune {
	scanner := bufio.NewReader(os.Stdin)

	for {
		choice, _, _ := scanner.ReadRune()

		if choice == ' ' || choice == '\n' {
			continue
		}

		return choice
	}
}

func getNumber() float32 {
	var (
		number float32
		err    error
	)

	for {
		_, err = fmt.Scanf("%f", &number)
		if err != nil {
			fmt.Print("Введите число, такое как 2.5, -1.78E8, или 3: ")
			_, err = fmt.Scanf("%d", &number)
			continue
		}

		return number
	}
}
