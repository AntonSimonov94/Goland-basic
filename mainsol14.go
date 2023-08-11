/*
Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.

Sample Input:
38012732
3
Sample Output:
801272
*/

package main
import (
	"fmt" 
)
func main () {
	var enterNum string
	var deletNum string
	var outNum string

	_, _ = fmt.Scan(&enterNum)
	_, _ = fmt.Scan(&deletNum)

	for _, elem := range enterNum {
		if deletNum != string(elem) {
			outNum += string(elem)
		}
	}
	fmt.Println(outNum)
}