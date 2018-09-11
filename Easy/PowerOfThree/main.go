package main

import (
	"fmt"
)

func main(){
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(0))
	fmt.Println(isPowerOfThree(9))
	fmt.Println(isPowerOfThree(45))
}
func isPowerOfThree(n int) bool {
    if n ==0 {
		return false
	}
	if n ==1 {
		return true
	}
	if n%3 !=0 {
		return false
	}
	n/=3
	return isPowerOfThree(n)
}
func isPowerOfThree2(n int) bool {
    x := 1
    for x < n {
        x = x*3
    }
    
    return x == n
}