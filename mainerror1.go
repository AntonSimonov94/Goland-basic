/*
Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!

Не забудьте считать два числа которые необходимо поделить (передать в функцию)!

Sample Input:
10 5
Sample Output:
2
*/


package main

import (
    "errors"
    "fmt"
)

func main() {
var  a, b int
fmt.Scan(&a)
fmt.Scan(&b)
	res, err := divide(a, b)
	fmt.Print(err != nil)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(res)
		
	}
}


func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	} else {
		return a / b, nil
	}
	}