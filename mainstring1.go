package main
import (
	"fmt"
	"bufio"
	"os"
	"unicode"
	"unicode/utf8"
)

func main () {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//first := rune(text[0])
	end := []rune(text)
	if (unicode.IsUpper(end[0])) && (string(end[utf8.RuneCountInString(text)-3]) == ".") {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}	
}
//Быть или не быть. utf8.RuneCountInString(text)-2