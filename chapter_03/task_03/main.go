package main

import "fmt"

// Упражнение по программированию 3.3
//
// Напишите программу, которая выдаст предупредительный звуковой сигнал, а затем выводит следующий текст
// Напуганная внезапным звуком, Вика вскрикнула:
// "Во имя всех звезд, что это было!"
func main() {
	fmt.Println("\aНапуганная внезапным звуком, Вика вскрикнула:")
	fmt.Println("\"Во имя всех звезд, что это было!\"")
}