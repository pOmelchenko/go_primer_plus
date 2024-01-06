package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Упражнение по программированию 8.2
//
// Напишите программу, которая читает ввод как поток символов, пока не встретит <code>EOF</code>. Программа
// должна выводить каждый введенный символ и его десятичный код ASCII. Следует отметить, что в кодировке ASCII символы,
// предшествующие пробелу, являются непечатаемыми. Трактуйте их особым образом. Если непечатаемым символом является
// символ новой строки или символ табуляции, выводите, соответственно, <code>\\n</code> или <code>\\t</code>.
// В противном случае воспользуйтесь нотацией управляющих символов. Например, ASCII-код 1 — это комбинация <CTRL+A>,
// которую можно отобразить как <code>^A</code>. Обратите внимание, что ASCII-код символа <code>A</code> представляет
// собой значение <Ctrl+A> плюс 64. Аналогичная зависимость имеется и для других непечатаемых символов. Выводите по 10
// пар в строке, кроме случая, когда встречается символ новой строки.
//
// Операционная система может иметь специальные интерпретации для некоторых управляющих символов и не допускать
// их попадания в программу.
func main() {
	var (
		count = 1
	)

	scanner := bufio.NewReader(os.Stdin)

	for {
		input, err := scanner.ReadByte()
		if err == io.EOF {
			break
		}

		switch input {
		case 0:
			fmt.Print("^@")
		case 1:
			fmt.Print("^A")
		case 2:
			fmt.Print("^B")
		case 3:
			fmt.Print("^C")
		//case 4: // EOF
		//	return
		case 5:
			fmt.Print("^E")
		case 6:
			fmt.Print("^F")
		case 7:
			fmt.Print("^G")
		case 8:
			fmt.Print("^H")
		case 9:
			fmt.Print("^I")
		case 10:
			fmt.Print("^J")
		case 11:
			fmt.Print("^K")
		case 12:
			fmt.Print("^L")
		case 13:
			fmt.Print("^M")
		case 14:
			fmt.Print("^N")
		case 15:
			fmt.Print("^O")
		case 16:
			fmt.Print("^P")
		case 17:
			fmt.Print("^Q")
		case 18:
			fmt.Print("^R")
		case 19:
			fmt.Print("^S")
		case 20:
			fmt.Print("^T")
		case 21:
			fmt.Print("^U")
		case 22:
			fmt.Print("^V")
		case 23:
			fmt.Print("^W")
		case 24:
			fmt.Print("^X")
		case 25:
			fmt.Print("^Y")
		case 26:
			fmt.Print("^Z")
		case 27:
			fmt.Print("^[")
		case 28:
			fmt.Print("^\\")
		case 29:
			fmt.Print("^]")
		case 30:
			fmt.Print("^^")
		case 31:
			fmt.Print("^_")
		case 32:
			fmt.Print(" ")
		default:
			fmt.Print(string(input))
		}

		fmt.Printf(" - %d ", input)

		if count%10 == 0 {
			fmt.Println()
		}

		count++
	}

	fmt.Println()
}
