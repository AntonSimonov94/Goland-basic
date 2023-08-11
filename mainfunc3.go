package main
import "fmt"
func main () {
var a, b int
fmt.Scan(&a)
n, sum := sumInt(a, b)
fmt.Printf("%d, %d", n, sum)
}
func sumInt(a ... int) (n int, s int) {
	sum := 0
	arr := []int{}
	for _, elem := range a {
		sum += elem
		arr = append(arr, elem)
	}
	return len(arr), sum
}
