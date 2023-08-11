package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	a := "%^80"
	b := "hhhhh20&&&&nd"

	fmt.Print(adding(a,b))
	//fmt.Print(strconv.ParseInt(a,10,0))
}

func adding(a,b string) int {
	var aRed, bRed string
	for _, aRes := range a {
		if unicode.IsDigit(aRes) {
			aRed += string(aRes)
		}
		
	}
	for _, bRes := range b {
		if unicode.IsDigit(bRes) {
			bRed += string(bRes)
		}
		
	}
    aInt, err1 := strconv.Atoi(aRed)
    bInt, err2 := strconv.Atoi(bRed)
	if err2 != nil && err1 != nil {
		panic(err1)
    } else {
		return aInt + bInt 
	}
}