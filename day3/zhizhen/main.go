package main

import "fmt"

func arrayModify(a1 *[3]int) {
	var a2 = [3]int{1, 2, 3}
	fmt.Printf("%T   %v\n", a1, a1)
	fmt.Printf("%T   %v\n", a2, a2)
}

func main() {
	a1 := [3]int{1, 2, 3}
	arrayModify(&a1)
	fmt.Printf("%v", a1)
}
