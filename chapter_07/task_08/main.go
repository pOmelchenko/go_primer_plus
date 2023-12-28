package main

import (
	"fmt"
)

const (
	NormalWorkTime      float32 = 40
	OvertimeCoefficient float32 = 1.5
)

// Упражнение по программированию 7.8
//
// Измените предположение а) в упражнении 7 так, чтобы программа предоставляла меню с тарифными ставками. Для
// выбора тарифной ставки используйте оператор <code>switch</code>. После запуска программы вывод должен быть подобным
// показанному ниже:
// <pre>
// ***********************************************************************
// Введите число, соответствующее желаемой тарифной ставке или действию:<br>
// 1) $8.75/ч                     2) $9.33/ч<br>
// 3) $10.00/ч                    4) $11.20/ч<br>
// 5) Выход
// ***********************************************************************
// </pre>
// если выбранный вариант с 1 по 4, программа должна запрашивать количество отработанных часов. Программа должна
// повторяться до тех пор, пока не будет вызван вариант 5. При вводе чего-то отличного от вариантов 1-5 программа
// должна напомнить пользователю допустимые варианты для ввода и снова ожидать ввод. Для различных тарифных и налоговых
// ставок применяйте константы <code>#define</code>
func main() {
	var (
		menu           int
		hours          float32
		normalWorkTime float32
		earned         float32
		tax            float32
		err            error
	)

	for {
		fmt.Printf("***********************************************************************\n")
		fmt.Printf("1) $8.75/ч                     2) $9.33/ч\n")
		fmt.Printf("3) $10.00/ч                    4) $11.20/ч\n")
		fmt.Printf("5) Выход\n")
		fmt.Printf("***********************************************************************\n")
		_, err = fmt.Scan(&menu)
		if err != nil {
			return
		}

		switch menu {
		case 1:
			normalWorkTime = 8.75
		case 2:
			normalWorkTime = 9.33
		case 3:
			normalWorkTime = 10.00
		case 4:
			normalWorkTime = 11.20
		case 5:
			return
		default:
			fmt.Printf("Попробуйте еще раз\n")
		}

		fmt.Printf("Введите количество отработанных часов: ")
		_, err = fmt.Scan(&hours)
		if err != nil {
			return
		}

		if hours <= 40 {
			earned = hours * normalWorkTime
		} else {
			normalEarned := NormalWorkTime * normalWorkTime
			overtimeEarned := (hours - NormalWorkTime) * (normalWorkTime * OvertimeCoefficient)

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
}
