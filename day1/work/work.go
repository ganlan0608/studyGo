package main

import "fmt"

//字符串翻转操作
func str() {
	s1 := "hello橄榄"
	var s2 string

	for _, i := range s1 {
		s2 = string(i) + s2
	}
	println(s2)

}
func str2() {
	s1 := "hello"
	bytesArray := []byte(s1)
	bytesArray[0], bytesArray[len(bytesArray)-1] = bytesArray[len(bytesArray)-1], bytesArray[0]
	fmt.Printf("%s", bytesArray)
}

func main() {
	str2()

}
