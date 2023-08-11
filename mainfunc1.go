/*
Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:
4 5 6 7
Sample Output:
4
*/

package main
import "fmt"
func main () {
var x,y,z,q int
fmt.Scan(&x, &y, &z, &q)
result := minFunc(x,y,z,q)
fmt.Printf("%d", result)
}
func minFunc (a,b,c,d int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	} 
	if d < min {
		min = d
	}
	return min
}