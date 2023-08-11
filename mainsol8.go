/*
Количество минимумов
Найдите количество минимальных элементов в последовательности.

Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные

Выведите количество минимальных элементов последовательности.

Sample Input:
3
21
11
4
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
	min := arr[0]
	for i:= 0; i<len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			num = 1
		} else if arr[i] == min {
			num++
		}
	}
	fmt.Print(num)
}