package main

import (
    "errors"
    "fmt"
)

func main() {
var  a, b int
fmt.Scan(&a)
fmt.Scan(&b)
	res, err := divide(a, b)
	fmt.Print(err != nil)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(res)
		
	}
}


func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	} else {
		return a / b, nil
	}
	}