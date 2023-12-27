package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	NORMAL_WORK_TIME     = 40
	NORMAL_TARIFF_RATE   = 10
	OVERTIME_COEFFICIENT = 1.5
)

// Упражнение по программированию 7.7
//
// Напишите программу, которая запрашивает количество часов, отработанных за неделю, и выводит значение общей
// суммы начислений, налогов и чистой зарплаты. Исходите из перечисленных ниже утверждений.
// <ul>
// <li>Основная тарифная ставка заработной платы = $10.00/час</li>
// <li>Сверхурочные часы (превышающие 40 часов в неделю) = коэффициент 1,5</li>
// <li>Налоговая ставка:<ul>
// <li>15% с первых $300;</li>
// <li>20% со следующих $150;</li>
// <li>25% с остатка.</li>
// </ul>
// </li>
// </ul>
// Используйте константы <code>#define</code> и не беспокойтесь, если приведенный пример не соответствует действующему
// налогообложения.
func main() {
	var (
		input  rune
		hours  int
		earned float32
		tax    float32
		err    error
	)

	fmt.Print("Введите количество отработанных часов: ")

	reader := bufio.NewReader(os.Stdin)

	input, _, err = reader.ReadRune()
	if err != nil {
		log.Println("stdin: ", err)
		return
	}

	hours, err = strconv.Atoi(string(input))
	if err != nil {
		log.Println("convert data: ", err)
		return
	}

	if hours <= 40 {
		earned = float32(hours * NORMAL_TARIFF_RATE)
	} else {
		normalEarned := NORMAL_WORK_TIME * NORMAL_TARIFF_RATE
		overtimeEarned := (hours - NORMAL_WORK_TIME) * (NORMAL_TARIFF_RATE * OVERTIME_COEFFICIENT)

		earned = float32(normalEarned + overtimeEarned)
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
