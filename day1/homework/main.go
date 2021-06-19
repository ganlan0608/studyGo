package main

import "fmt"

func primes(start, end int) {
	if start == 1 {
		start++
	}
	for i := start; i <= end; i++ {
		prime(i)
	}

}
func prime(v int) {
	var x bool
	for i := 2; i < v/2; i++ {
		if v%i == 0 {
			x = true
			break
		}
	}
	if x != true {
		fmt.Printf("%d是一个素数\n", v)
	}

}

func nineNine() {
	var num []string

	for i := 1; i <= 9; i++ {
		num = []string{}
		for x := 1; x <= i; x++ {
			numStr := fmt.Sprintf("%d*%d=%d", x, i, x*i)
			num = append(num, numStr)
		}
		fmt.Printf("%v\n", num)

	}
}

func main() {
	nineNine()
	primes(1, 10)

}
