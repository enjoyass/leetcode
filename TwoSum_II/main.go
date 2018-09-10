package main

import (
	"fmt"
)

func main() {
	nums := []int{2,7,11,15}
	target := 9
	re:=twoSum(nums,target)
	fmt.Println(re)
}

func twoSum(numbers []int, target int) []int {
	l:= 0
	r:=len(numbers)-1
	for l<=r{
		temp:=numbers[l]+numbers[r];
		if temp==target{
			return []int{l+1,r+1}
		}else if temp<target{
			l++
		}else{
			r--
		}
	}
	return []int{}
}