package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NormalWorkTime      = 40
	OvertimeCoefficient = 1.5
)

// Упражнение по программированию 8.7
//
// Модифицируйте упражнение по программированию 8 из главы 7 так, чтобы пункты меню помечались буквами, а не
// цифрами; для прекращения ввода используйте букву <code>q</code> вместо <code>5</code>.
func main() {
	var (
		hours          int
		choice         rune
		normalWorkTime float32
		earned         float32
		tax            float32
	)

	for choice != 'q' {
		choice = getChoice()
		switch choice {
		case 'a':
			normalWorkTime = 8.75
		case 'b':
			normalWorkTime = 9.33
		case 'c':
			normalWorkTime = 10.00
		case 'd':
			normalWorkTime = 11.20
		default:
			fmt.Println("Попробуйте еще раз")
		}

		fmt.Print("Введите количество отработанных часов: ")
		hours = getInt()

		if hours <= 40 {
			earned = float32(hours) * normalWorkTime
		} else {
			normalEarned := NormalWorkTime * normalWorkTime
			overtimeEarned := float32(hours-NormalWorkTime) * (normalWorkTime * OvertimeCoefficient)

			earned = normalEarned + overtimeEarned
		}

		if earned <= 300 {
			tax = earned * 15 / 100
		} else if earned <= 450 {
			secondPart := earned - 300

			tax = 300 * 15 / 100
			tax += secondPart * 20 / 100
		} else {
			thirdPart := earned - 300 - 150

			tax = 300 * 15 / 100
			tax += 150 * 20 / 100
			tax += thirdPart * 25 / 100
		}

		fmt.Printf("Начислено: %.2f\nНалог %.2f\nЧистый доход: %.2f\n", earned, tax, earned-tax)
	}

	fmt.Println("Программа завершена.")
}

func getChoice() rune {
	fmt.Println("***********************************************************************")
	fmt.Println("a) $8.75/ч                     b) $9.33/ч")
	fmt.Println("c) $10.00/ч                    d) $11.20/ч")
	fmt.Println("q) Выход")
	fmt.Println("***********************************************************************")

	ch := getFirst()

	for {
		if ch != 'a' || ch != 'b' && ch != 'c' && ch != 'd' && ch != 'q' {
			fmt.Println("Выберите a, b, c, d или q.")
			ch = getFirst()
		} else {
			return ch
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

func getInt() int {
	var (
		input int
		//ch    rune
	)

	for {
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			fmt.Println(" не является целочисленным.")
			fmt.Print("Введите целое число, такое как 25, -178 или 3: ")
		}
	}
}
