package main

import (
	"fmt"
)

func main(){
	fmt.Println(generate(6))
}
func generate(numRows int) [][]int {
	if numRows <=0 {
		return [][]int{}
	}
    if numRows == 1 {
	   	return [][]int{[]int{1}}
    }
	result :=[][]int{[]int{1}}
	for i:=1;i<numRows;i++ {
		temp:=[]int{}
		for j:=0;j<=len(result[i-1]);j++ {
			if j==0 || j==len(result[i-1]){
				temp=append(temp,1)
			}else{
				tempData :=result[i-1][j]+result[i-1][j-1]
				temp=append(temp,tempData)
			}
		}
		result=append(result,temp)
	}
	return result
}