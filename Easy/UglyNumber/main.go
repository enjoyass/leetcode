package main

import (
	"fmt"
)

func main(){
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
}
func isUgly(num int) bool {
	if num==1 {
		return true
	}

	if num %2!=0 && num%3!=0 && num%5!=0 {
		return false
	}
	
	if num %2 ==0 {
		num/=2
	}
	if num %3 ==0 {
		num/=3
	}
	if num %5 ==0 {
		num/=5
	}
	return isUgly(num)
}