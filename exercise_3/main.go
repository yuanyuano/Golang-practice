package main

import "fmt"

func main() {
	var a = [5]int{1, 3, 5, 7, 8}
	/*数组求和*/
	var add = 0
	for _, v1 := range a {
		add = add + v1
	}
	fmt.Printf("Sum: %d", add)
	fmt.Println()
	/*找出和为8的两个元素*/
	var j int
	for i, v2 := range a {
		for j = i + 1; j < 5 && j > i; j++ {
			if v2+a[j] == 8 {
				fmt.Printf("(%d %d)\t", i, j)
			}
		}
	}
}
