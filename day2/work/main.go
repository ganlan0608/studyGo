package main

import "fmt"

func arraySum(arrayA [5]int) int {
	var sum int
	for _, i := range arrayA {
		sum = sum + i
	}
	arrayA = [5]int{1: 2}

	return sum
}

func arraySum8(arrayA [5]int) {
	sumAnd := 8
	for index, v := range arrayA {
		for i := index + 1; i < len(arrayA); i++ {
			if v+arrayA[i] == sumAnd {
				println(index, i)
			}

		}
	}
}

func main() {
	arrayA := [5]int{1, 3, 5, 7, 9}
	sum := arraySum(arrayA)
	println(sum)
	fmt.Printf("%v\n", arrayA)
	arraySum8(arrayA)
}
