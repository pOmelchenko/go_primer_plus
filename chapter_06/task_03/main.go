package main

import "fmt"

const START_LETTER = 'F'

const END_LETTER = 'A'

// Упражнение по программированию 6.3
//
// Воспользуйтесь вложенными циклами, чтобы написать программу, которая выводит следующую последовательность
// символов:
// <pre>
// F<br>
// FE<br>
// FED<br>
// FEDC<br>
// FEDCB<br>
// FEDCBA
// </pre>
// Примечание: если в вашей системе не используется <code>ASCII</code> или какая-то другая кодировка, в которой буквы
// представлены в числовом порядке, то для инициалиазции символьного массива буквами алфавита вы можете применить
// следующее объявление:
// <pre>
// chars lets[27] = "ABSCDEFGHIJKLMNOPQRSTUVWXYZ";
// </pre>
// Затем для выбора конкретных букв можно использовать индексы массива, на пример, <code>lets[0]</code> для 'А' и т.д.
func main() {
	var (
		row int
		col int
	)

	for row = 0; row <= START_LETTER-END_LETTER; row++ {
		for col = 0; col <= row; col++ {
			fmt.Printf("%c", START_LETTER-col)
		}
		fmt.Println()
	}
}
