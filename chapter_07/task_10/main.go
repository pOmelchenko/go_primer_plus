package main

import "fmt"

const (
	TaxCategorySingle              = 17850
	TaxCategoryMarriage            = 23900
	TaxCategoryMarriageAndJoint    = 29750
	TaxCategoryMarriageAndSeparate = 14875
)

// Упражнение по программированию 7.10
//
// В 1988 году шкала федеральных налоговых ставок Соединенных Штатов была самой простой за всё прошедшее время.
// Она содержала четыре категории, каждая из которых включала две ставки. Ниже приведены самые общие данные (суммы в
// долларах представляют собой доход, облагаемый налогом)
// <table>
//
//	<thead>
//	  <tr><th>Категория</th><th>Налог</th></tr>
//	</thead>
//	<tbody>
//	  <tr>
//	    <td>Одинокий</td>
//	    <td>15% с первых $17850 плюс 28% от суммы, превышающей указанную</td>
//	  </tr>
//	  <tr>
//	    <td>Глава семейства</td>
//	    <td>15% с первых $23900 плюс 28% от суммы, превышающей указанную</td>
//	  </tr>
//	  <tr>
//	    <td>Состоит в браке, совместное ведение хозяйства</td>
//	    <td>15% с первых $29750 плюс 28% от суммы, превышающей указанную</td>
//	  </tr>
//	  <tr>
//	    <td>Состоит в браке, раздельное ведение хозяйства</td>
//	    <td>15% с первых $14875 плюс 28% от суммы, превышающей указанную</td>
//	  </tr>
//	</tbody>
//
// </table>
// Например, одинокий работник, получающий облагаемый налогом доход в $20000, платит налоги в сумме
// 0.15х$17850 + 0.28х($20000 - $17850). Напишите программу, которая позволяет пользователю указать категорию и
// облагаемый налогом доход, после чего вычисляет сумму налога. Используйте цикл, чтобы пользователь мог вводить
// разные варианты налогообложения.
func main() {
	var (
		income             float32
		taxAmount          float32
		incomeOverTheLimit float32
	)

	for {
		fmt.Print("Введите облагаемый налогом доход: ")
		if _, err := fmt.Scan(&income); err != nil {
			fmt.Println("Не верный формат дохода")
			return
		}
		var menu int
		fmt.Print("Выберите один из вариантов налогообложения: ")

		if _, err := fmt.Scan(&menu); err != nil {
			fmt.Println("Не верный вариант меню")
			return
		}

		switch menu {
		case 1:
			fmt.Println("Выбран вариант: Одинокий")

			if income <= TaxCategorySingle {
				taxAmount = income * 0.15
			} else {
				taxAmount = TaxCategorySingle * 0.15
				incomeOverTheLimit = income - TaxCategorySingle
			}
		case 2:
			fmt.Println("Выбран вариант: Глава семейства")

			if income <= TaxCategoryMarriage {
				taxAmount = income * 0.15
			} else {
				taxAmount = TaxCategoryMarriage * 0.15
				incomeOverTheLimit = income - TaxCategoryMarriage
			}
		case 3:
			fmt.Println("Выбран вариант: Состоит в браке, совместное ведение хозяйства")

			if income <= TaxCategoryMarriageAndJoint {
				taxAmount = income * 0.15
			} else {
				taxAmount = TaxCategoryMarriageAndJoint * 0.15
				incomeOverTheLimit = income - TaxCategoryMarriageAndJoint
			}
		case 4:
			fmt.Println("Выбран вариант: Состоит в браке, раздельное ведение хозяйства")

			if income <= TaxCategoryMarriageAndSeparate {
				taxAmount = income * 0.15
			} else {
				taxAmount = TaxCategoryMarriageAndSeparate * 0.15
				incomeOverTheLimit = income - TaxCategoryMarriageAndSeparate
			}
		}
		taxAmount += incomeOverTheLimit * 0.28

		fmt.Printf("Размер налога: %0.2f\n", taxAmount)

		fmt.Print("Выберите один из вариантов налогообложения: ")
	}
}
