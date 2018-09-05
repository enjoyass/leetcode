package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(-12321))
}

func isPalindrome(x int) bool {
	str :=strconv.Itoa(x)
	var newStr []byte
	for i:=0;i<len(str);i++ {
		newStr=append(newStr,str[len(str)-i-1])
	}
	for i:=0;i<len(str);i++ {
		if str[i] != newStr[i] {
			return false
		}
	}
	return true
}