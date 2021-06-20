package main

import (
	"fmt"
	"strings"
)

//返回
func strCounts(str string, b string) int {

	return strings.Count(strings.ToLower(str), b)
}

func main() {
	coins := 50
	coinsMap := map[string]int{"e": 1, "i": 2, "o": 3, "u": 4}
	users := [...]string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Perter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	nameCoins := make(map[string]int)

	for i := range coinsMap {
		for _, j := range users {
			coin := strCounts(j, i) * coinsMap[i] //每个字母在一个人身上分配到多少金币
			nameCoins[j] += coin
			coins -= coin
		}
	}
	for k, v := range nameCoins {
		fmt.Printf("%s分到%d个金币\n", k, v)

	}
	fmt.Printf("总共还有%d个金币", coins)
}
