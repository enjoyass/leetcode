package main

import (
	"fmt"
)

func main() {
	haystack := "abbbbbaabbaabaabbbaaaaabbabbbabbbbbaababaabbaabbbbbababaababbbbaaabbbbabaabaaaabbbbabbbaabbbaabbaaabaabaaaaaaaa"
	needle := "abbbaababbbabbbabbbbbabaaaaaaabaabaabbbbaabab"
	fmt.Println(strStr(haystack,needle))
}
func strStr(haystack string, needle string) int {
	if len(needle)==0 {
		return 0
	}
	for i:=0;i<=len(haystack)-len(needle);i=i+1{
		str :=string([]byte(haystack)[i:i+len(needle)])
		if len(str) == len(needle) {
			fmt.Println(str ,needle)
			if str == needle {
				fmt.Println(i)
				return i
			}
		}else{
			return -1
		}
		
	}
    return -1
}