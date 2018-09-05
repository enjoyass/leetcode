package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7,11,15,6,3,14}
	target := 9
	re:=twoSum(nums,target)
	fmt.Println(re)
}

func twoSum(nums []int, target int) []int {
  mapData :=make(map[int]int)
  numsLen:=len(nums)
  var result []int
  for i:=0;i<numsLen;i++ {
	mapData[nums[i]]=i
  }
  fmt.Println(mapData)
  for i:=0;i<numsLen;i++ {
	complement := target - nums[i];
	if val,ok:=mapData[complement]; ok && val!=i {
		
		result = append(result,i,mapData[complement])
		delete(mapData,complement)
		delete(mapData,nums[i])
	}
  }
  return result
}