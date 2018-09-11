package main

import (
	"fmt"
)

func main(){
	s := "anagram" 
	t := "nagaram"
	fmt.Println(isAnagram(s,t))
}
func isAnagram(s string, t string) bool {
    if s=="" && t=="" {
		return true
	}
	if len(s)!=len(t) {
		return false
	}
	sByte:= []rune(s)
	tByte:=[]rune(t)

	for i:=0;i<len(sByte);i++{
		for j:=i;j<len(sByte);j++ {
			if sByte[i]<sByte[j] {
				sByte[i],sByte[j]=sByte[j],sByte[i]
			}
			if tByte[i]<tByte[j] {
				tByte[i],tByte[j]=tByte[j],tByte[i]
			}
		}
	}
	if string(sByte) == string(tByte) {
		return true
	}
	return false
}