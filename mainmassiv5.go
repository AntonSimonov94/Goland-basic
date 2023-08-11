package main
import "fmt"

func main () {
	var n int
	fmt.Scan(&n)
	arr := make([]int,n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}
	num := 0
	for i:=0; i < n; i++{
		if arr[i] > 0 {
			num += 1
		}
	}
	fmt.Printf("%d", num)
}