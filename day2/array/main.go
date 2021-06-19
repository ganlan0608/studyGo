package main

import "fmt"

func main() {
	var a [5]int
	var b = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var c [10]int
	a = [5]int{1, 2, 3, 4, 5}
	c = [10]int{5: 10, 7: 111}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	println(a[2])
	for i := 0; i < len(a); i++ {
		print(a[i])
	}
	println()
	for _, v := range a {
		print(v)
	}
}
