package main

import (
	"fmt"
)

func main(){
	fmt.Println(Other(2))
}
func isHappy(n int) bool {
	m := map[int]bool{}
    for n != 1 && !m[n] {
        m[n] = true
        var t int
        for ; n != 0; n /= 10 {
            d := n % 10
            t += d * d
        }
        n = t
    }
    return n == 1
}
func Other(n int)bool{
	m := map[int]bool{}
	for n!=1&& !m[n] {
		m[n] = true
		var temp int
		for ;n!=0;n/=10{
			d:=n%10
			temp+=d*d
		}
		n=temp
		fmt.Println(n)
	}
	return n == 1
}