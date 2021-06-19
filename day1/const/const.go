package main

//iota 每次自增1
const (
	_ = iota
	a
	b
	c
	d
)

func main() {
	println(a, b, c, d)
}
