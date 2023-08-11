/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").

Sample Input:
топот
Sample Output:
Палиндром
*/

package main
import (
	"fmt"
	"bufio"
	"os"
)

func main () {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//first := rune(text[0])
	first := []rune(text)
	n := len(first)-3
	for i := range first {
		if (i == n) || (i > n) {
			fmt.Println("Полиндром")
			break
		}
		if first[i] == first[n] {
			n--
		} else {
			fmt.Println("Нет")
			break
		}

	} 
}
