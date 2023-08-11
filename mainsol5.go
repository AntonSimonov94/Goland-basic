package main
import "fmt"
func main () {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if ((a + b) > c) && ((a + c) > b) && ((c + b) > a) {
		fmt.Print("Существует")
	} else { fmt.Print("Несуществует") }
}