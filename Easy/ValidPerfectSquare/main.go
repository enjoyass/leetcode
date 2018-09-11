package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPerfectSquare(16))
	fmt.Println(isPerfectSquare(14))
}

func isPerfectSquare(num int) bool {
    r :=num

	for r*r>num {
		r=(r+num/r)/2
	}
	return r*r==num
}