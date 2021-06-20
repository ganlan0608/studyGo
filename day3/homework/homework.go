package main

import (
	"fmt"
)

func main() {

	var student struct {
		id    int
		name  string
		age   int
		score int
	}

	a := student
	a.id = 1
	a.name = "橄榄"
	a.age = 10
	a.score = 22
	fmt.Printf("%v\n", a)

	s := make(map[int]struct {
		id    int
		name  string
		age   int
		score int
	}, 10)
	s[1] = a

	fmt.Printf("map&struct:%v\n", s)
	fmt.Printf("%v\n", s[1])
	students := make(map[int]map[string]string, 10)
	students[1] = make(map[string]string, 3)
	students[1]["age"] = "18"
	students[1]["score"] = "88"
	students[1]["name"] = "橄榄"

	students[2] = make(map[string]string, 3)
	students[2]["age"] = "22"
	students[2]["score"] = "77"
	students[2]["name"] = "小明"

	fmt.Printf("%v", students[2])

}
