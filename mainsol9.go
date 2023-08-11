package main
import "fmt"
func main () {
	var n int
	fmt.Scan(&n)
	del := 10000000000
	result := sumArr(makeArr(n, del))
		fmt.Printf("%d ", sumArr(makeArr(result, del)))
	}
	func makeArr(num int, del int) []int {
		arr := []int{}
		for ; del > 0 ; {
			res := num / del
			 if (res != 0) {
				arr = append(arr, res)
				num = num - res*del
			} else if num < 10  {
				arr = append(arr, num % 10)
				break
			} 
				del /=10
			}
			return arr
	}
	func sumArr (arr []int ) int {
		sum := 0
		for i:= 0; i < len(arr); i ++ {
			sum += arr[i] 
		}
		return sum
	}
