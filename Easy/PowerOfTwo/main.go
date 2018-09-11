package main

import (
	"fmt"
)

func main(){
	fmt.Println(isPowerOfTwo2(1))
	fmt.Println(isPowerOfTwo2(8))
	fmt.Println(isPowerOfTwo2(218))
}
func isPowerOfTwo(n int) bool {
    if n ==0 {
		return false
	}
	if n ==1 {
		return true
	}
	if n%2 !=0 {
		return false
	}
	n/=2
	return isPowerOfTwo(n)
}

func isPowerOfTwo2(n int) bool {
    x := 1
    for x < n {
        x = x*2
    }
    
    return x == n
}