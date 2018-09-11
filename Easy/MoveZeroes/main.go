package main

import (
	"fmt"
)

func main(){
	moveZeroes2([]int{0,1,0,3,12})
}
func moveZeroes(nums []int)  {
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			if nums[i]==0 && nums[j] !=0 {
				nums[i],nums[j]=nums[j],nums[i]
			}
		}
	}
    fmt.Println(nums)
}
func moveZeroes2(nums []int)  {
    nz := 0
    i := 0
    for i < len(nums) {
        if nums[i] != 0 {
            nums[nz], nums[i] = nums[i], nums[nz]
            nz++
		}
		fmt.Println(nums)
        i++
	}
}