package main

import (
	"fmt"
)

func main(){
	// fmt.Println(intersection([]int{4,9,5},[]int{9,4,9,8,4}))
	// fmt.Println(intersection([]int{1,2,2,1},[]int{2,2}))
	// fmt.Println(intersection([]int{1,2},[]int{1,1}))
	fmt.Println(intersection([]int{3,1,2},[]int{1}))
}
func intersection(nums1 []int, nums2 []int) []int {
    s := []int{}

	if len(nums1) != 0 && len(nums2) != 0 {
		m := make(map[int]bool)

		for _, v := range nums1 {
			if _, ok := m[v]; !ok {
				m[v] = false
			}
		}

		for _, v := range nums2 {
			if b, ok := m[v]; ok && b == false {
				m[v] = true
				s = append(s, v)
			}
		}
	}
	return s
}