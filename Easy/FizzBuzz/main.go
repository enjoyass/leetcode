package main

import (
	"strconv"
	"fmt"
)

func main(){
	fmt.Println(fizzBuzz(15))
}
func fizzBuzz(n int) []string {
	var i int =1
	var result []string
	for i<=n {
		var tmp string
		if i%15==0 {
			tmp="FizzBuzz"
		}else if i%3 ==0 {
			tmp="Fizz"
		}else if i%5 ==0 {
			tmp="Buzz"
		}else{
			tmp=strconv.Itoa(i)
		}
		result=append(result,tmp)
		i++
	}
	return result
}