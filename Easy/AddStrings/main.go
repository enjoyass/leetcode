package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(addStrings("1","1"))
}

func addStrings(num1 string, num2 string) string {
    res := ""
	i,j,carry := len(num1) - 1, len(num2) - 1, 0
	for i >= 0 || j >= 0 {
		var a,b int
		if i>=0 {
			a,_=strconv.Atoi(string(num1[i]))
		}else{
			a=0
		}
		if j>=0 {
			b,_=strconv.Atoi(string(num2[j]))
		}else{
			b=0
		}
		
		sum:= a + b + carry;
		res=strconv.Itoa(sum%10)+res
		carry = sum / 10
		i--
		j--
	}
	if carry != 0 {
		return  "1" + res
	}
	return res
}