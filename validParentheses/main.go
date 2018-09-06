package main

import (
	"fmt"
)

func main(){
	fmt.Println(isValid("(("))
}
func isValid(s string) bool {
	if len(s)%2!=0 {
		return false
	}
	var tempData []byte
	for i :=0;i<len(s);i++ {
		if s[i] == '(' {
			tempData=append(tempData,')')
		}else if s[i]=='[' {
			tempData=append(tempData,']')
		}else if s[i]=='{' {
			tempData=append(tempData,'}')
		}else if len(tempData) ==0 ||tempData[len(tempData)-1] != s[i] {
			return false
		}else if tempData[len(tempData)-1] == s[i]{
			tempData=tempData[:len(tempData)-1]
		}
	}
	return len(tempData)==0
}