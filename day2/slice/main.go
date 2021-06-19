package main

import "fmt"

func main() {

	var a = []int{1, 2, 3}
	//fmt.Printf("%d\n", a)
	var b []int
	b = make([]int, 2, 9)
	//b = a
	//a[0] = 5
	//fmt.Printf("%d\n", b)
	copy(b, a)
	a[0] = 100
	fmt.Println(b)
	fmt.Println(a)

	//fmt.Printf("%d\n", cap(b))
}
