package main

import (
	"fmt"
)

func main(){
	fmt.Println(-1<<32)
	fmt.Println(thirdMax([]int{2,2,3,1}))
	fmt.Println(thirdMax([]int{2,1}))
}
func thirdMax(nums []int) int {
  min,mid,max:=-1<<32,-1<<32,-1<<32
  for _,val :=range nums {
	  if val >max {
		min=mid
		mid=max
		max=val
	  }else if val < max && val > mid {
		min = mid
        mid = val
	  }else if val < mid && val > min {
        min = val
      }
  }
  if min == -1<<32 {
	return int(max)
  } 
  return int(min)
}