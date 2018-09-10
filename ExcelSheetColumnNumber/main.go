package main

import (
	"fmt"
)

func main(){
	str :="AB"
	fmt.Println(titleToNumber(str))
}
func titleToNumber(s string) int {
	num:=0
	for i:=len(s)-1;i>=0;i--{
		if i>0 {
			tmepdata:=1
			n:=i
			for n>0 {
				tmepdata=tmepdata*26
				n--
			}
			num+=tmepdata*(int(s[len(s)-i-1]-'A')+1)
		}
	}
	num+=(int(s[len(s)-1]-'A')+1)
	return num
}