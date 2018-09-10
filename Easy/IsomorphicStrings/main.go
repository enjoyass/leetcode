package main

import (
	"fmt"
)

func main(){
	s := "abc"
	t := "bcd"
	fmt.Println(isIsomorphic2(s,t))
}
func isIsomorphic(s string, t string) bool {
	if s=="" || t=="" {
		return true
	}
	if len(s)!=len(t) {
		return false
	}
	m := make(map[byte]byte)
	d := make(map[byte]bool)
	for i := 0; i < len(s); i++ {
		if val, ok := m[s[i]]; ok {
			if val != t[i] {
				return false
			}
			continue
		}
		if _, ok := d[t[i]]; ok {
			return false
		}
		m[s[i]] = t[i]
		d[t[i]] = true
	}
	return true
}

func isIsomorphic2(s string, t string) bool {
	if len(s) != len(s){
		return false
	}
	Arr1 := make([]int, 256)
	Arr2 := make([]int, 256)

	for i:=0; i<len(s); i++{
		if Arr1[s[i]] != Arr2[t[i]]{
			return false
		}
		Arr1[s[i]] = i+1;
		Arr2[t[i]] = i+1;
		
	}
	return true
}