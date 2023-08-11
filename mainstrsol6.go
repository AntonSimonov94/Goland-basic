package main

import (
	"fmt"
	"math"
)

func main() {
	var k, p, v int
	fmt.Scan(&k,&p,&v)
	fmt.Print(T(k, p, v))
}

func T(k, p, v int) (float64) {
	result := 6 / W(k, p, v)
	return result
}

func W(k, p, v int) float64 {
	return math.Sqrt(k / M(p, v))
}

func M(b, c int) int {
	return b * c
}
//LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO