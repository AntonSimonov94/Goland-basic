package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
		if (n > 6) && (n < 21) || (n % 10 > 4) || (n % 10 == 0){
			fmt.Printf("%d korov", n)
		} else if (n % 10 == 1){
			fmt.Printf("%d korova", n)
		} else {fmt.Printf("%d korovy", n)}
}
		
		
	