/*
 На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:
ihgewlqlkot
Sample Output:
hello
*/


package main
import (
	"fmt"
	"bufio"
	"os"
)

func main () {
	var text2 string
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	for i, num := range text1 {
		if (i % 2 == 1) {
			text2 += string(num)
		}
	}
	fmt.Println(text2)
}
