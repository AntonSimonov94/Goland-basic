package main
import (
	"fmt" 
)
func main () {
	var enterNum string
	var deletNum string
	var outNum string

	_, _ = fmt.Scan(&enterNum)
	_, _ = fmt.Scan(&deletNum)

	for _, elem := range enterNum {
		if deletNum != string(elem) {
			outNum += string(elem)
		}
	}
	fmt.Println(outNum)
}