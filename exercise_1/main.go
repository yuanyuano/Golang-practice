package main

import "fmt"

func main() {
	var a = 1
	var b = 1.56
	var c = true
	var d = "Hello World!"
	var s1 = "hello我的名字是袁媛"
	var count = 0
	for _, c1 := range s1 {
		if len(string(c1)) > 1 {
			count++
		}
	}
	fmt.Printf("a的值为：%d  \t类型为：%T\n", a, a)
	fmt.Printf("b的值为：%f  \t类型为：%T\n", b, b)
	fmt.Printf("c的值为：%t  \t类型为：%T\n", c, c)
	fmt.Printf("d的值为：%s  \t类型为：%T\n", d, d)
	fmt.Println(s1)
	fmt.Printf("汉字的数量为：%d", count)
}
