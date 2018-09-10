package main

import (
	"fmt"
)

func main(){
	fmt.Println(rob2([]int{2,1,1,2}))
}
func rob(nums []int) int {
    if len(nums) == 0 { return 0}
    if len(nums) == 1 { return nums[0]}

    x,y,z := 0, nums[0], nums[1]
    for i:= 2; i < len(nums); i++ {
		fmt.Println(x,y,z)
        oldZ := z
        z = max(nums[i] + x, nums[i] + y)
        
        x = y
		y = oldZ
		fmt.Println(x,y,z)
    }
    
    return max(max(x,y),z)
}

func rob2(nums []int)int{
	a:=0
	b:=0
	for i:=0;i<len(nums);i++{
		if i % 2 == 0 {
			a = max(a + nums[i], b);
		} else {
			b = max(a, b + nums[i]);
		}
		fmt.Println(a,b)
	}
	return max(a,b)
}
func max(x,y int) int {
    if x < y {return y}
    return x
}