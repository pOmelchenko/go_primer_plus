package main

import "fmt"

// Упражнение по программированию 3.6
//
// Масса одной молекулы воды приблизительно составляет 3.0х10^-23 грамма. Кварта воды весит примерно 950 грамм.
// Напишите программу, которая предлагает ввести значение объема воды в квартах и отображает количество молекул воды в
// этом объеме.
func main() {
	var molecule_of_water_in_gram float32 = 3.0e-23
	var volume_of_water float32

	fmt.Print("Введите вес воды в граммах: ")
	fmt.Scan(&volume_of_water)
	fmt.Printf(
		"Количество молекул в заданном весе (%.2f) равно: %.0f",
		volume_of_water,
		volume_of_water/molecule_of_water_in_gram)
}
