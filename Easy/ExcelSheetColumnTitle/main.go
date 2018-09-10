package main

import (
	"fmt"
)

func main(){
	fmt.Println(convertToTitle(28))
}
func convertToTitle(n int) string {
	tempdata:=[]byte{}
	for n > 0 {
        tempdata = append(tempdata, 'A' + byte((n - 1) % 26))
        n = (n - 1) / 26
    }
	for i:=0;i<len(tempdata)/2;i++{
		tempdata[i],tempdata[len(tempdata)-i-1]=tempdata[len(tempdata)-i-1],tempdata[i]
	}
    return string(tempdata)
}
