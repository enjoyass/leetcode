package main

import (
	"fmt"
)

func main(){
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3},1))
}
func containsNearbyDuplicate(nums []int, k int) bool {

	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		idx, ok := mp[nums[i]]
		if(ok && i - idx <= k){
			return true
		}
		mp[nums[i]] = i
	}
	return false
}