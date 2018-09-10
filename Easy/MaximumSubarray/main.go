package main

import (
	"fmt"
)


func main(){
 nums :=[]int{4,-5,-4,3,10,-9,4,1}
 fmt.Println(maxSubArray(nums))
}


func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxSubArray(nums []int) int {
    suffix := make([]int, len(nums))
    results := make([]int, len(nums))
    
    if len(nums) == 0 {
        return 0
    }
		
    suffix[0] = nums[0]
    results[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        cursuffix := max(suffix[i-1] + nums[i], nums[i])
		suffix[i] = cursuffix
		results[i] = max(results[i-1], cursuffix)
    }
    return results[len(nums) - 1]
}
