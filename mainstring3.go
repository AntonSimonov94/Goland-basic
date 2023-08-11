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
	//text2, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	
	//first := []rune(text1)
	//second := []rune(text2)
	fmt.Println(strings.Trim(text1, "\n"))
	//fmt.Println(len(strings.Trim(text2, "\n")))
	//fmt.Print(strings.Index(string(text1), strings.Trim(text2, "\n")))
	fmt.Print(strings.Index(string(text1), text2.Text()))
}
