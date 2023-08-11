package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	h := n / 3600
	m := (n - h*3600) / 60
	fmt.Printf("It is %d hours %d minutes.", h, m)
}