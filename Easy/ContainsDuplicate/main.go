package main

import (
	"fmt"
)

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,1}))
}
func containsDuplicate(nums []int) bool {
   m :=make(map[int]int)
   for i:=0;i<len(nums);i++{
	   if _,ok:=m[nums[i]];ok{
		   return true
	   }
	   m[nums[i]]=i
   }
   return false
}