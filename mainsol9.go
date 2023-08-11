/*
Цифровой корень
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 . 

По данному числу определите его цифровой корень.

Входные данные

Вводится одно натуральное число n, не превышающее 
10 7.

Выходные данные
Вывести цифровой корень числа n.

Sample Input:
3456
Sample Output:
9
*/


package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	del := 10000000000
	result := sumArr(makeArr(n, del))
		fmt.Printf("%d ", sumArr(makeArr(result, del)))
	}
	func makeArr(num int, del int) []int {
		arr := []int{}
		for ; del > 0 ; {
			res := num / del
			 if (res != 0) {
				arr = append(arr, res)
				num = num - res*del
			} else if num < 10  {
				arr = append(arr, num % 10)
				break
			} 
				del /=10
			}
			return arr
	}
	func sumArr (arr []int ) int {
		sum := 0
		for i:= 0; i < len(arr); i ++ {
			sum += arr[i] 
		}
		return sum
	}
