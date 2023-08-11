/*
Количество нулей
По данным числам, определите количество чисел, которые равны нулю.  

Входные данные 
Вводится натуральное число N, а затем N чисел.

Выходные данные 
Выведите количество чисел, которые равны нулю.

Sample Input:
5
1
8
100
0
12
Sample Output:
1
*/

package main
import "fmt"
func main () {
	var n, m int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n ; i++ {
		fmt.Scan(&m)
		arr[i] = m
	}
	num := 0
	for i:= 0; i<len(arr); i++ {
		if arr[i] == 0 {
			num++
		}
	}
	fmt.Print(num)
}