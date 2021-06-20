package main

import "fmt"

func main() {
	type student struct {
		id    int
		name  string
		age   int
		score int
	}
	students := map[int]student{}
	students[1] = student{
		id:    1,
		name:  "ganlan",
		age:   18,
		score: 99,
	}
	students[2] = student{
		id:    2,
		name:  "sunc",
		age:   19,
		score: 22,
	}
	fmt.Printf("%v", students)
}
