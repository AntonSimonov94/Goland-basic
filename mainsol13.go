/*
Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:
8
Sample Output:
6
*/

package main
import (
	"fmt" 
)
func main () {
	var n int
	fmt.Scan(&n)
	arr := []int{1,1}
	if (arr[0] == n) {
		fmt.Printf("%d", 1)
	} else if (arr[1] == n) {
		fmt.Printf("%d", 2)
	} else {
	for i:= 2;  ; i++ {
			arr = append(arr, arr[i-2]+arr[i-1])
			if arr[i] == n {
				fmt.Printf("%d", i+1)
				break
			}
			if arr[i] > n {
				fmt.Printf("%d", -1)
				break
			}
		}
	}
}
		
		
	