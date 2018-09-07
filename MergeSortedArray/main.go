package main

import (
	"fmt"
	
)

func main() {
	nums1 := []int{1,4,6}
	nums2 :=[]int{2,3,5}
	merge(nums1,3,nums2,3)
}
func merge(nums1 []int, m int, nums2 []int, n int)  {
	nums1=nums1[:m]
	nums1 =append(nums1,nums2[:n]...)
	for i:=0;i<len(nums1);i++ {
		for j:=i+1;j<len(nums1);j++ {
			if nums1[i] > nums1[j] {
				nums1[i],nums1[j]=nums1[j],nums1[i]
			}
		}
	}
	fmt.Println(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int)  {
    if n == 0 {
        return
    }
    c1 := m - 1
    c2 := n - 1
    c0 := m + n - 1
    for c2 >= 0 {
        if c1 >= 0 && nums1[c1] > nums2[c2] {
            nums1[c0] = nums1[c1]
            c1--
        } else {
            nums1[c0] = nums2[c2]
            c2--
        }
        c0--
    }
}