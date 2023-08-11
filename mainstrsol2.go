/*
Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков.

Выходные данные

Вывести строку, которая получится после добавления символов '*'.

Sample Input:
LItBeoFLcSGBOFQxMHoIuDDWcqcVgkcRoAeocXO
Sample Output:
L*I*t*B*e*o*F*L*c*S*G*B*O*F*Q*x*M*H*o*I*u*D*D*W*c*q*c*V*g*k*c*R*o*A*e*o*c*X*O
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)
func main () {
	var text2 string
	text1, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text1 = text1[:len(text1)-2]
	for i, let := range text1 {
		if i != (len([]rune(text1))-1) {
			text2 += string(let) + "*"
		} else { 
			text2 += string(let)
			break 
		}
		
	}
	fmt.Println(text2)
}
