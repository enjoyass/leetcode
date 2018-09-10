package main

import (
	"fmt"
)

func main(){
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	k:=0
    for i:=1;i<n;i++{
		if isPrime(i){
			k++
		}
	} 
	return k
}

func isPrime(value int) bool {
    if value <= 3 {
        return value >= 2
    }
    if value%2 == 0 || value%3 == 0 {
        return false
    }
    for i := 5; i*i <= value; i += 6 {
        if value % i == 0 || value %(i+2) == 0 {
            return false
        }
    }
    return true
}