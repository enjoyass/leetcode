package main

import (
	"fmt"
)

func main(){
	fmt.Println(5==getSum(3,2))
}
func getSum(a int, b int) int {
    sum := a;
	carry := b; 
	for carry!=0{
		temp := (sum ^ carry);
		carry = (sum & carry) << 1;
		sum = temp;
	}
	return sum;
}