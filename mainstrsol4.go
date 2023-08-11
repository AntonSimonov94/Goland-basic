package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"os"
)
func main () {
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text1 = text1[:len(text1)-2]
	text2 := ""
	a := []rune(text1)
	for _, elem := range a {
		newNum, _ := strconv.ParseInt(string(elem), 10, 64)
		newElem := int(math.Pow(float64(newNum),2))
		text2 += strconv.Itoa(int(newElem))
	}
	fmt.Println(text2)
}
//LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO