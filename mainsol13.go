package main
import (
	"fmt" 
)
func main () {
	var n int
	fmt.Scan(&n)
	arr := []int{1,1}
	if (arr[0] == n) {
		fmt.Printf("%d", 1)
	} else if (arr[1] == n) {
		fmt.Printf("%d", 2)
	} else {
	for i:= 2;  ; i++ {
			arr = append(arr, arr[i-2]+arr[i-1])
			if arr[i] == n {
				fmt.Printf("%d", i+1)
				break
			}
			if arr[i] > n {
				fmt.Printf("%d", -1)
				break
			}
		}
	}
}
		
		
	