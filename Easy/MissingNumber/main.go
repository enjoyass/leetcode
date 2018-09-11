package main

import (
	"fmt"
)

func main(){
	nums:=[]int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(nums))
}
func missingNumber(nums []int) int {
	maxValue :=0
	sum1:=0
	sum2:=0
	for i:=0;i<len(nums);i++ {
		sum1+=i
		sum2+=nums[i]
		if maxValue<nums[i] {
			maxValue=nums[i]
		}
	}
	if sum1==sum2 {
		return maxValue+1
	}
	return sum1+maxValue-sum2
}