package main

import "fmt"

const (
	PRICE_OF_ARTICHOKES = 2.05
	PRICE_OF_BEETS      = 1.15
	PRICE_OF_CARROT     = 1.09

	DISCOUNT_VALUE = 5
)

// Упражнение по программированию 7.11
//
// Компания ABC Main Order Grocery продает артишоки по цене $2.05 за фунт, свеклу по $1.15 за фунт и морковь по
// $1.09 за фунт. До добавления затрат на доставку компания предоставляет скидку 5% на заказы на сумму до $100 и выше.
// Затраты составляют $6.50 за доставку и обработку заказа весом в 5 фунтов или менее, $14.00 за обработку и доставку
// заказа весом от 5 до 20 фунтов и $14.00 плюс $0.50 за каждый фунт для доставки заказа весом, превышающим 20 фунтов.
// Напишите программу, которая использует оператор <code>switch</code> в цикле так, что в ответ на ввод <code>a</code>
// пользователь получает возможность указать желаемый вес артишоков в фунтах; в ответ на ввод <code>b</code> — вес
// свеклы в фунтах; в ответ на ввод <code>c</code> — вес моркови в фунтах; а в ответ на <code>q</code> — завершить
// процесс заказа. Программа должна вести учет сумм нарастающим итогом. То есть если пользователь вводит 4 фунта свеклы
// и позже вводит еще 5 фунтов свеклы, программа должна сообщать о заказе 9 фунтов свеклы. Затем программа должна
// вычислить общие затраты, скидку, если есть, расходы на доставку и полную сумму заказа. Далее программа должна
// отобразить всю информацию о покупке: стоимость фунта товара, количество заказанных фунтов, стоимость каждого
// заказанного вида овощей, общую стоимость заказа, скидку (если есть), затраты на доставку и итоговую сумму заказа с
// учетом всех затрат.
func main() {
	var (
		menu                        string
		weight_of_artichokes        float32
		weight_of_beets             float32
		weight_of_carrot            float32
		weight                      float32
		summary_weight              float32
		summary_price_of_artichokes float32
		summary_price_of_beets      float32
		summary_price_of_carrot     float32
		summary_price               float32
		discount                    float32
		delivery_price              float32
		final_price                 float32
	)

	fmt.Println("Добро пожаловать в ABC Main Order Grocery!")
	fmt.Println()
	fmt.Println("У нас вы можете приобрести:")
	fmt.Printf("a) артишоки — $%.2f за фунт\n", PRICE_OF_ARTICHOKES)
	fmt.Printf("b) свеклу — $%.2f за фунт\n", PRICE_OF_BEETS)
	fmt.Printf("c) морковь — $%.2f за фунт\n", PRICE_OF_CARROT)
	fmt.Println()
	fmt.Print("Выберите товар который вы хотите добавить в корзину (или введите q чтобы закончить заказ): ")

loop:
	for {
		if _, err := fmt.Scan(&menu); err != nil {
			return
		}

		switch menu {
		case "a":
			fmt.Printf("Укажите вес артишоков: ")
		case "b":
			fmt.Printf("Укажите вес свеклы: ")
		case "c":
			fmt.Printf("Укажите вес моркови: ")
		case "q":
			break loop
		default:
			fmt.Printf("Такого товара не найдено, попробуйте еще раз: ")
		}

		if _, err := fmt.Scan(&menu); err != nil {
			return
		}

		switch menu {
		case "a":
			weight_of_artichokes += weight
			summary_price_of_artichokes += weight_of_artichokes * PRICE_OF_ARTICHOKES
			fmt.Printf("Добавлено %0.2f фунта(-ов) артишоков\n", weight)
		case "b":
			weight_of_beets += weight
			summary_price_of_beets += weight_of_beets * PRICE_OF_BEETS
			fmt.Printf("Добавлено %0.2f фунта(-ов) свеклы\n", weight)
		case "c":
			weight_of_carrot += weight
			summary_price_of_carrot += weight_of_carrot * PRICE_OF_CARROT
			fmt.Printf("Добавлено %0.2f фунта(-ов) моркови\n", weight)
		}

		fmt.Printf("Если вы хотите добавить что-то еще, выберите товар: ")
	}

	fmt.Println("Чек")
	fmt.Printf("Артишоков\n\tвесом: %0.2f фунтов\n\tна сумму: $%0.2f\n\tпо $%0.2f за фунт\n",
		weight_of_artichokes,
		summary_price_of_artichokes,
		PRICE_OF_ARTICHOKES)
	fmt.Printf("Свеклы\n\tвесом: %0.2f фунтов\n\tна сумму: $%0.2f\n\tпо $%0.2f за фунт\n",
		weight_of_beets,
		summary_price_of_beets,
		PRICE_OF_BEETS)
	fmt.Printf("Моркови\n\tвесом: %0.2f фунтов\n\tна сумму: $%0.2f\n\tпо $%0.2f за фунт\n",
		weight_of_carrot,
		summary_price_of_carrot,
		PRICE_OF_CARROT)

	summary_weight = weight_of_artichokes + weight_of_beets + weight_of_carrot
	summary_price = summary_price_of_artichokes + summary_price_of_beets + summary_price_of_carrot

	fmt.Printf("ИТОГО\n\tобщий вес: %0.2f фунтов\n\tна сумму: $%0.2f\n",
		summary_weight,
		summary_price)

	if summary_price <= 100 {
		discount = summary_price * DISCOUNT_VALUE / 100
		fmt.Printf("\tскидка %d%% суммой: $%0.2f\n", DISCOUNT_VALUE, discount)
	}

	if summary_weight <= 5 {
		delivery_price = 6.50
	} else if summary_weight > 5 && summary_weight <= 20 {
		delivery_price = 14.00
	} else {
		delivery_price = 0.50 * summary_weight
	}

	fmt.Printf("\tстоимость доставки: %.2f\n", delivery_price)

	final_price = summary_price - discount + delivery_price

	fmt.Printf("\nК ОПЛАТЕ:\n\t$%.2f", final_price)
}
