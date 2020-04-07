package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, value := range users {
		for _, v := range strings.Split(value, "") {
			switch {
			case v == "e" || v == "E":
				coins--
			case v == "i" || v == "I":
				coins -= 2
			case v == "o" || v == "O":
				coins -= 3
			case v == "u" || v == "U":
				coins -= 4
			}
		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
