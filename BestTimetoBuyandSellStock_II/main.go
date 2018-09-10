package main
import (
	"fmt"
)
func main(){
	prices := []int{1,2,3,4,5,6,7}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maxValue:=0
	for i:=0;i<len(prices)-1;i++{
		if prices[i+1] >prices[i] {
			maxValue+=(prices[i+1]-prices[i])
		}
		
	}
    return maxValue
}