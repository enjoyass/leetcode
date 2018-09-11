package main

import (
	"fmt"
)

func main(){
	fmt.Println(reverseString("A man, a plan, a canal: Panama"))
}
func reverseString(s string) string {
	byteArr:=[]byte(s)
	for i:=0;i<len(s)/2;i++{
		byteArr[i],byteArr[len(s)-i-1]=byteArr[len(s)-i-1],byteArr[i]
	}
	return string(byteArr)
}