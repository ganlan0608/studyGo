package main

import "fmt"

func main() {
	var a [2]int32
	a = [2]int32{1, 2}
	var b [3][2]int32
	b = [3][2]int32{
		a,
		{3, 4},
	}

	fmt.Println(a)
	fmt.Println(b)

}
