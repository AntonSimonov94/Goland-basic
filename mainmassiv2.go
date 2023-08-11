package main
import "fmt"

func main () {
	var N int
	fmt.Scan(&N)
	arr := make([]int,N,N)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}

	fmt.Printf("%v", arr[3])
}
