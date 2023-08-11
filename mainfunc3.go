/*
Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1
*/

package main
import "fmt"
func main () {
var a, b int
fmt.Scan(&a)
n, sum := sumInt(a, b)
fmt.Printf("%d, %d", n, sum)
}
func sumInt(a ... int) (n int, s int) {
	sum := 0
	arr := []int{}
	for _, elem := range a {
		sum += elem
		arr = append(arr, elem)
	}
	return len(arr), sum
}
