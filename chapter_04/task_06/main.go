package main

import "fmt"

// Упражнение по программированию 4.6
//
// Напишите программу, которая запрашивает имя пользователя и его фамилию. Сделайте так, чтобы она выводил
// введенные имена в одной строке и количество символов в каждом слова в следующей строке. Выровняйте количество
// символов по окончанию соответствующего имени, как показано ниже:
// <pre>Иван Петров</pre>
// <pre>   4      6</pre>
// Затем сделайте так, чтобы программ выводила ту же самую информацию. но с количеством символов, выровненным по началу
// каждого слова:
// <pre>Иван Петров</pre>
// <pre>4    6</pre>
func main() {
	var first_name string
	var second_name string

	fmt.Print("Введите имя и фамилию: ")
	fmt.Scan(&first_name, &second_name)
	fmt.Printf("%s %s\n", first_name, second_name)
	fmt.Printf("%*d %*d\n", len(first_name), len(first_name), len(second_name), len(second_name))
	fmt.Println()
	fmt.Printf("%s %s\n", first_name, second_name)
	fmt.Printf("%-*d %-*d\n", len(first_name), len(first_name), len(second_name), len(second_name))
}
