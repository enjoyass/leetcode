package main

import (
	"fmt"
)

func main(){
	nums := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(nums,2))
}
func removeElement(nums []int, val int) int {
	if len(nums)==0{
		return 0
	}
	index:=0
	for i:=0;i<len(nums);i++{
		if nums[i]!=val{
			nums[index]=nums[i]
			index++
		}
	}
	nums=nums[:index]
    return len(nums)
}