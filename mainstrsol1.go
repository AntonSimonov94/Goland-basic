/*
На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:
6 8
Sample Output:
10
*/

package main
import (
	"fmt"
	"math"
)
func main () {
	var a, b int 
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Print(int(pifagor(a,b)))
}

func pifagor(a,b int) (float64) {
	return math.Sqrt(math.Pow(float64(a),2)+math.Pow(float64(b),2))
}