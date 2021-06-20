package main

import (
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

func main() {
	aa()
	b()
	aa()

}
