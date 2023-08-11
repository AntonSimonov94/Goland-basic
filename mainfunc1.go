package main
import "fmt"
func main () {
var x,y,z,q int
fmt.Scan(&x, &y, &z, &q)
result := minFunc(x,y,z,q)
fmt.Printf("%d", result)
}
func minFunc (a,b,c,d int) int {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	} 
	if d < min {
		min = d
	}
	return min
}