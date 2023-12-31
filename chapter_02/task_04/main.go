package main

import "fmt"

// Упражнение по программированию 2.4
//
// Напишите программу, которая производит следующий вывод:
// Он веселый молодец!
// Он веселый молодец!
// Он веселый молодец!
// Никто не может это отрицать!
// Вдобавок к функции <code>main()</code> в программе должны использоваться две определенные пользователем функции:
// <code>jolly()</code>, которая выводит сообщение "Он веселый молодец!" один раз, и <code>deny()</code>, выводящая
// сообщение в последней строке.
func main() {
	jolly()
	jolly()
	jolly()
	deny()
}

func jolly() {
	fmt.Println("Он веселый молодец!")
}

func deny() {
	fmt.Println("Никто не может это отрицать!")
}
