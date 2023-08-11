package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	e := n % 10
	d := (n / 10) % 10
	s := n /100
	num := e*100 + d*10 + s
	fmt.Printf("%d", num)
}