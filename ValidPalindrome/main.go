package main

import (
	"fmt"
)

func main(){
	s:="A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	var baseByte []byte
	var reveByte []byte
	for i:=0;i<len(s);i++ {
		if s[i] >='0' && s[i] <='9' {
			baseByte=append(baseByte,s[i])
			reveByte=append(reveByte,s[i])
		}else if s[i] >='a' && s[i] <='z' {
			baseByte=append(baseByte,s[i])
			reveByte=append(reveByte,s[i])
		}else if s[i] >='A' && s[i] <='Z' {
			baseByte=append(baseByte,s[i]+32)
			reveByte=append(reveByte,s[i]+32)
		}
	}
	for i:=0;i<len(reveByte)/2;i++{
		reveByte[i],reveByte[len(reveByte)-i-1]=reveByte[len(reveByte)-i-1],reveByte[i]
	}
	if string(baseByte)==string(reveByte) {
		return true
	}
    return false
}