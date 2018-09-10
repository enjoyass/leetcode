package main

import (
	"fmt"
)

func main() {
	digits:=[]int{9}
	fmt.Println(plusOne(digits))
}
func plusOne(digits []int) []int {
    for i:=len(digits)-1;i>=0;i--{
		digits[i]=digits[i]+1
		if digits[i] < 10 {
			break
		}else{
			digits[i]=digits[i]%2
			if i==0 {
				digits=append([]int{1},digits...)
			}
		}
	}
	fmt.Println(digits)
	return digits
}