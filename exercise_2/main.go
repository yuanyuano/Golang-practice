//9*9乘法表
package main

import "fmt"

func main() {
	var a, b int
	for a = 1; a <= 9; a++ {
		for b = 1; b <= a; b++ {
			fmt.Printf("%d*%d=%d\t", b, a, b*a)
		}
		fmt.Println()
	}
}
