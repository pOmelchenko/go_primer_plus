package main

import "fmt"

// Упражнение по программированию 4.2
//
// Напишите программу, которая запрашивает имя и выполняет с ним следующие действия.
//
// Выводит его заключенным в двойные кавычки.
// Выводит его в поле шириной 20 символов, при этом все поле заключается в кавычки, а имя выравнивается по правому краю поля.
// Выводит его с левого края поля ширирой 20 символов, при этом все поле заключается в кавычки.
// Выводит его в поле шириной, на три символа превышающего длину имени.
func main() {
	var name string
	fmt.Print("Введите имя: ")
	fmt.Scan(&name)

	// Выводит его заключенным в двойные кавычки.
	fmt.Printf("\"%s\"", name)
	fmt.Println()

	// Выводит его в поле шириной 20 символов, при этом все поле заключается в кавычки, а имя выравнивается по правому краю поля.
	fmt.Printf("\"%20s\"", name)
	fmt.Println()

	// Выводит его с левого края поля ширирой 20 символов, при этом все поле заключается в кавычки.
	fmt.Printf("\"%-20s\"", name)
	fmt.Println()

	// Выводит его с левого края поля ширирой 20 символов, при этом все поле заключается в кавычки.
	fmt.Printf("\"%*s\"", len(name)+3, name)
	fmt.Println()
}
