package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(wordPattern("abba","dog cat cat dog"))
}
func wordPattern(pattern string, str string) bool {
	if len(str) == 0 {
		return false
	}
	strSlice:=strings.Split(str," ")
	if len(strSlice)!=len(pattern) {
		return false
	}
	appearance := make(map[string]bool)
	mapping := make(map[byte]string)

	for i := 0; i < len(pattern); i++ {
		ch, str := pattern[i], strSlice[i]
		val, ok := mapping[ch]
		if ok {
			if val == str {
				continue
			}
			return false
		}

		appeared, ok := appearance[str]
		if ok && appeared {
			return false
		}

		mapping[ch] = str
		appearance[str] = true
	}
	return true
}