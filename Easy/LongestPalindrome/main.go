package main

import (
	"fmt"
)

func main(){
	fmt.Println(longestPalindrome("abccccdd"))
	// fmt.Println(longestPalindrome("AAAAAA"))
}

func longestPalindrome(s string) int {
	apha:=make([]int,26)
	Apha:=make([]int,26)
	var lpd int
	var more int
	for _,v:=range s {
		if v >='a' && v<='z'{
			apha[v-'a']++
		}else{
			Apha[v-'A']++
		}
		
	}
	for i:=0;i<26;i++{
		if apha[i] >0 {
			if apha[i] %2 == 0 {
				lpd+= apha[i]
			}else{
				lpd+= apha[i]-1
				more=1
			}
		}
		if Apha[i] >0 {
			if Apha[i] %2 == 0 {
				lpd+= Apha[i]
			}else{
				more=1
				lpd+= Apha[i]-1
			}
		}
	}
	return lpd+more
}

func longestPalindrome2(str string) int {
    if len(str) == 0 {
	return 0
    }

    mapS := make(map[rune]int)
    ans := 0
    for _,s := range str {
	mapS[s]++
    }
    for _, mps := range mapS {
	ans = ans + mps / 2 * 2;
            if (ans % 2 == 0 && mps % 2 == 1){
		ans++
	} 
    }

    return ans
}