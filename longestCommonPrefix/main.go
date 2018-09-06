package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	prfix := strs[0]
	for i:=1;i<len(strs);i++{
		for !strings.HasPrefix(strs[i],prfix) {
			prfix=string([]byte(prfix)[0:len(prfix)-1])
			if len(prfix) == 0 {
				return ""
			}
		}
	}
    return prfix
}
