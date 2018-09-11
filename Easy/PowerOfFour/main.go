package main

import (
	"fmt"
)

func main(){
	fmt.Println(isPowerOfFour(5))
	fmt.Println(isPowerOfFour(16))
}

func isPowerOfFour(n int) bool {
    x := 1
    for x < n {
        x = x*4
    }
    
    return x == n
}