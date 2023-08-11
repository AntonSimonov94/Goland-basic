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
