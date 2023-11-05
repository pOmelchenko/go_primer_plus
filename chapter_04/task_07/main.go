package main

import "fmt"

// Упражнение по программированию 4.7
//
// Напишите программу, которая присваивает переменной типа <code>double</code> значение 1.0/3.0 и переменной
// типа <code>float</code> значение 1.0/3.0. Отобразите каждый результат три раза: в первом случае с четырьмя цифрами
// справа от десятичной точки, во втором случае с двенадцатью цифрами и в третьем случае с шестнадцатью цифрами.
// Включите также в программу заголовочный файл <code>float.h</code> и выведите значение <code>FLT_DIG</code> и
// <code>DBL_DIG</code>. Согласуются ли выведенные значения со значением 1.0/3.0?
func main() {
	var f_res float32 = 1.0 / 3.0
	var d_res float64 = 1.0 / 3.0

	fmt.Printf("float32:  %.4f %.12f %.16f\n", f_res, f_res, f_res)
	fmt.Printf("float64: %.4f %.12f %.16f\n", d_res, d_res, d_res)
}
