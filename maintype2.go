/*
Для решения данной задачи вам понадобится пакет strconv, возможно использовать пакеты strings или encoding/csv, или даже bufio - вы не ограничены в выборе способа решения задачи. В решениях мы поделимся своими способами решения этой задачи, предлагаем вам сделать то же самое.

В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".

На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).

Sample Input:
1 149,6088607594936;1 179,0666666666666
Sample Output:
0.9750
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	text1, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil && err != io.EOF {
		fmt.Print(err)
	}
	newString := strings.Split(text1[:len(text1)-2], ";")
		if len(newString) == 2 {
			num1, err1 := changeStr(newString[0])
			num2, err2 := changeStr(newString[1])
			if err1 != nil && err2 != nil {
				panic(errors.New("Ошибка"))
			}
			fmt.Printf("%.4f", num1/num2)
		} else {fmt.Println("Введено неверное число!")}
		
}

func changeStr(str string) (float64, error) {
 	str = strings.ReplaceAll(string(str), ",", ".")
 	str = strings.ReplaceAll(string(str), " ", "")
	newStr , err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
 	return newStr, nil
}
//1 149,6088607594936;1 179,0666666666666