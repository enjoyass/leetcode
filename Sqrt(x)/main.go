package main

import (
	"fmt"
)

func main() {
	if 2==mySqrt(8) {
		fmt.Println("SUCCESS")
	}else{
		fmt.Println("FAILED")
	}
}
func mySqrt(x int) int {
	r :=x

	for r*r>x {
		r=(r+x/r)/2
	}
	return r
}