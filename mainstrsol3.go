/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:
1112221112
Sample Output:
2
*/


package main

import (
	"bufio"
	"fmt"
	"os"
)
func main () {
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text1 = text1[:len(text1)-2]
	max := 0
	for _, elem := range text1 {
		if (int(elem) > max) {
			max = int(elem)
		}
	}
	fmt.Println(string(max))
}
