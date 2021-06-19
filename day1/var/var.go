package main

import "fmt"

var a string //声明全局变量

var b string = "ganlan999" //声明全局变量并且赋值，全局变量只能这样赋值
//批量赋值
var (
	name string
	age  int
)

func main() {
	var y = 100 //省略类型，系统自动推导
	fmt.Printf("y=%d\n", y)

	x := 200 //简单声明并赋值，系统自动推导类型  只能在函数内使用
	fmt.Printf("x=%v\n", x)
	a = "ganlan" //使用全局变量
	fmt.Printf("a=%s\n", a)
	fmt.Printf("b=%s\n", b)

	x, _ = fmt.Println(`c:\go`)

}
