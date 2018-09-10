package main

import (
	"fmt"
)

func main() {
	nums := []int {1,1,2,4,5,6,7,8,9}
	fmt.Println(removeDuplicates(nums))
}
func removeDuplicates(nums []int) int {
	if len(nums) ==0 {
		return 0
	}
	i:=0
	for j:=0;j<len(nums);j++{
		if nums[j]!=nums[i] {
			i++
			nums[i]=nums[j]
		}
	}
	nums=nums[:i+1]
    return len(nums)
}
