package main

import "fmt"

func main() {

	var a = []int{1, 2, 3}
	println(a)
	var b []int
	b = a
	a[0] = 5
	println(b)

	fmt.Printf("%d\n", b)
	fmt.Printf("%d\n", cap(b))
}
