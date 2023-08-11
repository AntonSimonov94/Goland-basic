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
