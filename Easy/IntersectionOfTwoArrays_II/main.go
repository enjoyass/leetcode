package main

import (
	"fmt"
)

func main(){
	// fmt.Println(intersection([]int{4,9,5},[]int{9,4,9,8,4}))
	// fmt.Println(intersection([]int{1,2,2,1},[]int{2,2}))
	// fmt.Println(intersection([]int{1,2},[]int{1,1}))
	fmt.Println(intersect([]int{3,1,2},[]int{1}))
}
func intersect(nums1 []int, nums2 []int) []int{
	var r []int
	m := make(map[int][]int)

	for _, v := range nums1 {
		if _, ok := m[v]; !ok {
			m[v] = make([]int, 2)
		}
		m[v][0]++
	}

	for _, v := range nums2 {
		if _, ok := m[v]; !ok {
			m[v] = make([]int, 2)
		}
		m[v][1]++
	}

	for k, v := range m {
		if v[0] != 0 && v[1] != 0 {
			if v[0] > v[1] {
				for i := 0; i < v[1]; i++ {
					r = append(r, k)
				}
			} else {
				for i := 0; i < v[0]; i++ {
					r = append(r, k)
				}
			}
		}
	}

	return r
}