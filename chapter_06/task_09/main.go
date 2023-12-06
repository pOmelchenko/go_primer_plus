package main

import "fmt"

// Упражнение по программированию 6.9
//
// Модифицируйте упражнение 8 так, чтобы программа использовала функцию для возврата результатов вычислений.
func main() {
	var (
		num_a float32
		num_b float32
		nums  = 2
	)

	fmt.Print("Введите пару числе чтобы произвести вычисления: ")

	for nums == 2 {
		nums, _ = fmt.Scanf("%f%f", &num_a, &num_b)
		fmt.Printf("(%f - %f) / (%f * %f) = %f\n",
			num_a,
			num_b,
			num_a,
			num_b,
			calculate(num_a, num_b))
		fmt.Println("Введите другую пару чисел, или не числовой символ чтобы прекратить работу:")
	}

	fmt.Println("Удачи!")
}

func calculate(num_a float32, num_b float32) float32 {
	return (num_a - num_b) / (num_a * num_b)
}
