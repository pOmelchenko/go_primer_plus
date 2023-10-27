package main

import "fmt"

// Упражнение по программированию 4.1
//
// Напишите программу, которая запрашивает имя и фамилию, а затем выводит их в формате <i>имя, фамилия</i>
func main() {
	var first_name string
	var second_name string

	fmt.Print("Введите имя и фамилию: ")

	fmt.Scan(&first_name, &second_name)

	fmt.Printf("%s %s", first_name, second_name)
}
