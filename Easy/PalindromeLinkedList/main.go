package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
func isPalindrome(head *ListNode) bool {
	if head==nil {
		return true
	}
	var value []int
	for head != nil {
		value=append(value,head.Val)
		head=head.Next
	}
	for i:=0;i<len(value)/2;i++{
		if value[i]!=value[len(value)-1-i] {
			return false
		}
	}
    return true
}
func main(){
	l4:=&ListNode{Val:1,Next:nil}
	l3:=&ListNode{Val:2,Next:l4}
	l2:=&ListNode{Val:1,Next:l3}
	l1:=&ListNode{Val:1,Next:l2}
	fmt.Println(isPalindrome(l1))
}