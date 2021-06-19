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

func arrayAndNum(arrayA [5]int, sumAnd int) (sumSlice [][]int) {
	for index, v := range arrayA {
		for i := index + 1; i < len(arrayA); i++ {
			if v+arrayA[i] == sumAnd {
				a := []int{index, i}
				sumSlice = append(sumSlice, a)
			}
		}
	}
	return sumSlice
}

func main() {
	arrayA := [5]int{1, 3, 5, 7, 9}
	sum := arraySum(arrayA)
	println(sum)
	x := arrayAndNum(arrayA, 10)
	fmt.Printf("%d\n", x)
}
