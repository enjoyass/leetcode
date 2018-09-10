package main

import (
	"fmt"
)

func main() {
    fmt.Println(romanToInt("LVIII"))
}
func romanToInt(s string) int {
    var mapData =map[rune]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    result :=0
    lastValue:=0
    for _,v :=range s {
        if lastValue < mapData[v] {
            result=result+mapData[v]-lastValue*2
        }else{
            result=result+mapData[v]
        }
        lastValue=mapData[v]
    }
    return result
}
