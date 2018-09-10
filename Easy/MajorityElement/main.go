package main

import (
	"fmt"
)

func main(){
	nums :=[]int{1,1,1,2,2,2,2}
	res:=majorityElement2(nums)
	if res==2{
		fmt.Println("SUCCESS")
	}else{
		fmt.Println("Failed")
	}
}
// slow
func majorityElement(nums []int) int {
	index:=0
	max:=0
    for i:=0;i<len(nums);i++{
		k:=0
		for j:=i+1;j<len(nums);j++{
			if nums[i]==nums[j] {
				k++
			}
		}
		if max < k{
			max=k
			index=i
		}
	}
	return nums[index]
}

// fast
func majorityElement2(nums []int) int {
    k := nums[0]
    count := 0
    for _,n := range nums{
        if count == 0{
            k = n 
            count ++
            continue
        }
        if k == n{
            count ++
        }else{
            count --
        }
    }
    return k
}