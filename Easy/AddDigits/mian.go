package main

import (
	"fmt"
)

func main(){
	fmt.Println(addDigits(38))
}
func addDigits(num int) int {
	if num <= 9 {
		return num
	}
	newNum:=0
    for num >9{
		newNum+=num%10
		num/=10
	}
	newNum=newNum+num
	return addDigits(newNum)
}