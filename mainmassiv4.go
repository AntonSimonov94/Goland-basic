package main
import "fmt"

func main () {
	var n int
	outArr := ""
	fmt.Scan(&n)
	arr := make([]int,n)
	for i := range arr {
		_, _ = fmt.Scan(&arr[i])
	}
	for i:=0; i < n; i++{
		if i % 2 == 0 {
			outArr += fmt.Sprintf("%d ",arr[i])
		}
	}
	fmt.Printf("%v", outArr)
}