package main

import (
	"fmt"
)

func main(){
	fmt.Println(canConstruct("bg","efjbdfbdgfjhhaiigfhbaejahgfbbbjagbddfgdiaigdadhcfcj"))
}
func canConstruct(ransomNote string, magazine string) bool {
    mp := make([]int, 26, 26)
    
    for i:=0; i < len(magazine); i++ {
        mp[int(magazine[i]) - int('a')] += 1  
    }
    
    for j:=0; j < len(ransomNote); j++ {
        mp[int(ransomNote[j]) - int('a')] -= 1  
        if mp[int(ransomNote[j]) - int('a')] < 0 {
                return false
        }
    }
    
    return true
}