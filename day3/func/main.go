package main

import (
	"fmt"
	"strings"
)

func intSum3(b int, a ...int) (sum int) {
	sum = b
	for _, arg := range a {
		sum = sum + arg
	}
	return
}

func testDefer() {
	defer println(1)
	defer println(2)
	defer println(3)
	println("func end")
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func a(name string) func() {
	return func() {
		println("say", name)
	}
}

func makeSuffixFunc(suffix string) func(string2 string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func aa() {
	println("hihi")

}

func b() {
	defer func() {
		if r := recover(); r != nil {

		}
	}()

	panic("err")

}
func f1() {
	var a1 []int
	a1 = make([]int, 0, 3)
	a1 = append(a1, 1, 2, 3)
	b1 := a1
	println("a1内存:", a1)
	fmt.Printf("a1值：%d\n", a1)
	println("b2内存:", b1)
	fmt.Printf("b1值：%d\n", b1)
	println("-----------------------")
	a1 = append(a1, 4, 5)
	println("a1内存:", a1)
	fmt.Printf("a1值：%d\n", a1)
	println("b2内存:", b1)
	fmt.Printf("b1值：%d\n", b1)

}

func main() {

}
