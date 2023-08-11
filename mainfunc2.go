package main
import "fmt"
func main () {
var n int
fmt.Scan(&n)
result := numFib(n)
fmt.Printf("%d", result)
}
func numFib(n int) int {
	arr := []int{1,1}
	for i:= 2; i <= n ; i++ {
			arr = append(arr, arr[i-2]+arr[i-1])
		}
	return arr[n-1]
}
