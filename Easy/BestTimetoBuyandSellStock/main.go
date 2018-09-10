package main

import (
	"fmt"
)

func main(){
	prices := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maxValue:=0
	for i:=0;i<len(prices);i++{
		for j:=i+1;j<len(prices);j++ {
			temp:=prices[j]-prices[i]
			if temp >maxValue {
				maxValue =temp
			}
		}
	}
    return maxValue
}