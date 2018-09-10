package main

import (
	"strings"
	"fmt"
)

func main(){
	str:="a "
	fmt.Println(lengthOfLastWord(str))
}
func lengthOfLastWord(s string) int {
	data:=strings.Split(s," ")
	for i:=len(data)-1;i>=0;i--{
		if len(data[i]) !=0 {
			return len(data[i])
		}
	}
	return 0
}