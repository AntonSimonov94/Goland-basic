/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:
fdsghdfgjsdDD1
Sample Output:
Ok
*/

package main
import (
	"fmt"
	"bufio"
	"os"
	//"strings"
	"unicode"
)

func main () {
	var text2 string
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	newString := text1[:len(text1)-2]
	if (len(newString)) >= 5 {
		for _, num := range newString {
			if (!unicode.IsLetter(num)) && (!unicode.IsDigit(num)) || (unicode.Is(unicode.Latin, num))  {
				text2 = "Wrong password"
				break
			} else {
				text2 = "Ok"
			}
		}
	} else { text2 = "Wrong password" }
	fmt.Print(text2)
}
