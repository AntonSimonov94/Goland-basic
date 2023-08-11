package main

import (
	"bufio"
	"fmt"
	"os"
)
func main () {
	var text2 string
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text1 = text1[:len(text1)-2]
	for i, let := range text1 {
		if i != (len([]rune(text1))-1) {
			text2 += string(let) + "*"
		} else { 
			text2 += string(let)
			break 
		}
		
	}
	fmt.Println(text2)
}
//LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO