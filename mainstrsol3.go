package main

import (
	"bufio"
	"fmt"
	"os"
)
func main () {
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text1 = text1[:len(text1)-2]
	max := 0
	for _, elem := range text1 {
		if (int(elem) > max) {
			max = int(elem)
		}
	}
	fmt.Println(string(max))
}

//LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO