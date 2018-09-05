package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println(reverse(-1534236469))
}
func reverse(x int) int {
	flag :=1 //1，正数，-1 负数

	 if x<0 && x>-2147483648 {
		flag=-1
		x=x*(-1)
	}else if x>0 && x <2147483647{
		flag =1
	}else {
		return 0
	}
	var newStr []byte
	str:=strconv.Itoa(x)
	for i:=0;i<len(str);i++ {
		newStr=append(newStr,str[len(str)-i-1])
	}
	for i:=0;i<len(newStr);i++ {
		if newStr[i]!='0' {
			newStr=newStr[i:]
			break
		}
	}
	result,_:=strconv.Atoi(string(newStr))
	if result > 2147483647 || result< -2147483648 {
		return 0
	}else {
		return flag*result
	}
   
}