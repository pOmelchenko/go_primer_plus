package main

import "fmt"

const FIRST_LETTER = 'A'

// Упражнение по программированию 6.5
//
// Напишите программу, которая предлагает пользователю ввести прописную букву. Воспользуйтесь вложенными
// циклами, чтобы написать программу, которая выводит фигуру в виде пирамиды, подобную приведенной ниже:
// <pre>
// &nbsp;&nbsp;&nbsp;&nbsp;A<br>
// &nbsp;&nbsp;&nbsp;ABA<br>
// &nbsp;&nbsp;ABCBA<br>
// &nbsp;ABCDCBA<br>
// ABCDEDCBA<br>
// </pre>
// Эта фигура должна расширяться вплоть до введенного символа. Например, представленная фигура стала результатом ввода
// символа <code>E</code>. Подсказка: для обработки строк воспользуйтесь внешним циклом. Для формирования строки
// применяйте три внутренних цикла: один для манипуляции пробелами, второй для вывода букв в порядке возрастания и
// трений для ввода букв в порядке убывания. Если в вашей системе не используется <code>ASCII</code> или подобная ей
// кодировка, в которой буквы представлены в числовом порядке, см. примечание в упражнении 3.
func main() {
	var (
		letter        int
		last_letter   int
		row           int
		rows          int
		space         int
		spaces        int
		letter_offset int
	)

	fmt.Printf("Введите букву от A до Z, чтобы программа построила пирамиду: ")
	fmt.Scanf("%c", &last_letter)

	rows = last_letter - FIRST_LETTER
	for row = 0; row <= rows; row++ {
		spaces = rows - row

		// spaces
		for space = 0; space < spaces; space++ {
			fmt.Print(" ")
		}

		// letters ++
		for letter_offset = 0; letter_offset <= row; letter_offset++ {
			fmt.Printf("%c", FIRST_LETTER+letter_offset)
		}

		// letters ++
		for letter = FIRST_LETTER + row - 1; letter >= FIRST_LETTER; letter-- {
			fmt.Printf("%c", letter)
		}

		fmt.Println()
	}
}
