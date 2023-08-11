package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main () {
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text2 := text1
	for _, num := range text1 {
		if (strings.Count(text1, string(num)) >= 2) {
			text2 = strings.ReplaceAll(text2, string(num), "")
		}
	}
	fmt.Print(text2)
}
//zaabcbd