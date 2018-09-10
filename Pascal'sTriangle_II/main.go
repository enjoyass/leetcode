package main

import (
	"fmt"
)

func main(){
	fmt.Println(generate(3))
}
func generate(numRows int) []int {
	if numRows <=0 {
		return []int{1}
	}
    if numRows == 1 {
	   	return []int{1,1}
    }
	result :=[]int{1}
	for i:=1;i<=numRows;i++ {
		temp:=[]int{}
		for j:=0;j<=len(result);j++ {
			if j==0 || j==len(result){
				temp=append(temp,1)
			}else{
				tempData :=result[j]+result[j-1]
				temp=append(temp,tempData)
			}
		}
		result=temp
	}
	return result
}