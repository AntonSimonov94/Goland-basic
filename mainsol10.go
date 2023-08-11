package main
import "fmt"
func main () {
	var a, b int
	fmt.Scan(&a, &b)
		for i := b ; i >= a; i-- {
			if i % 7 == 0 {
				fmt.Printf("%d ", i)
				break
			} else if i == a {
				fmt.Print("NO")
			}
		}
}
		
		
	