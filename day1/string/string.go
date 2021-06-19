package main

import (
	"Strings"
	"fmt"
	"strconv"
)

func main() {
	s1 := "ganlan"
	s2 := "好帅"
	fmt.Println(len(s1))
	fmt.Println(s1 + s2)
	s3 := fmt.Sprintf("%s------%s", s1, s2)
	fmt.Println(s3)
	s4 := "a:1,b:2,c:3,d:4"
	s5 := strings.Split(s4, ",")
	fmt.Println(s5)
	a1 := strings.Contains(s1, "gan")
	println(a1)
	a1 = strings.Contains(s1, "xx")
	println(a1)
	b2 := strings.HasPrefix(s1, "an")
	println(b2)
	c1 := strings.Index(s1, "n")
	println(c1)
	d3 := "hello橄榄"

	for i := 0; i < len(d3); i++ {
		fmt.Printf("%v(%c)\n", d3[i], d3[i])
	}

	for _, i := range d3 {
		fmt.Printf("%c\n", i)
	}
	s10 := "10"
	s11, _ := strconv.ParseInt(s10, 10, 32)
	println(s11)

	s12 := 11
	s13 := strconv.Itoa(s12)
	println(s13)

}
