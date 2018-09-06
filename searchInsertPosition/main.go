package main

import (
	"fmt"
)

func main(){
	nums :=[]int{1,3,5,6}
	fmt.Println(searchInsert(nums,5))
	fmt.Println(searchInsert(nums,2))
	fmt.Println(searchInsert(nums,7))
	fmt.Println(searchInsert(nums,0))
}

func searchInsert(nums []int, target int) int {
	for i,v:=range nums {
		  if v==target {
			  return i
		  }
		  if v>target {
			  return i
		  }
	  }
	  return len(nums)
  }