package main

import (
	"fmt"
)

func main(){
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}
func firstUniqChar(s string) int {
	freq:=make([]int,26)
	for i:=0;i<len(s);i++{
		freq [s[i] - 'a'] ++
	}
	for i:=0;i<len(s);i++{
		if freq[s[i] - 'a']==1 {
			return i
		}
	}
	return -1
}
