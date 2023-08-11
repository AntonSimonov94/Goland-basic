package main
import (
	"fmt" 
	"math")
func main () {
	var n int
	fmt.Scan(&n)
	for i := 0 ; i < n ; i++ {
		result := math.Pow(2, float64(i))
		if int(result) <= n {
			fmt.Printf("%d ", int(result))
		} else {break}
		
	}
}
		
		
	