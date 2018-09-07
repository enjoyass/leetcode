package main

import (
	"fmt"
)

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	var num int
	var arr []int 
	if n==1 {
		return 1
	}
	for i:=0;i<n;i++ {
		arr=append(arr,-1)
	}
	arr[0]=1;
	arr[1]=2;
	num =step(n-1,arr)
	return num
}

func step(n int, arr []int)int{
	if (arr[n]==-1) {
		arr[n] = step(n-1,arr)+step(n-2,arr)
		fmt.Println(arr)
	}
	return arr[n]
}