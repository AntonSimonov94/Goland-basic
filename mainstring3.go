/*
Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1

Sample Input:
awesome
es
Sample Output:
2
*/

package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main () {
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text2 := bufio.NewScanner(os.Stdin)
  	text2.Scan()
	fmt.Println(strings.Trim(text1, "\n"))
	fmt.Print(strings.Index(string(text1), text2.Text()))
}
