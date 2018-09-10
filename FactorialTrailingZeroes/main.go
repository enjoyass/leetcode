package main

import (
	"fmt"
)

func main(){
fmt.Println(trailingZeroes(200))
}
func trailingZeroes(n int) int {
	ret :=0
	pow:=5

    for pow*5>pow && n/pow >=1{
		fmt.Println(ret,pow)
		ret += n / pow
		pow *= 5
	}
	return ret
}