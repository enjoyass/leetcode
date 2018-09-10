package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil{
		return l2
	} 
    if l2 == nil{
		return l1
	} 

	if (l1.Val <l2.Val){
		l1.Next=mergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next=mergeTwoLists(l2.Next, l1)
		return l2
	}
}

func main(){
	lA3 :=&ListNode{Val:4,Next:nil}
	lA2 :=&ListNode{Val:2,Next:lA3}
	lA1 :=&ListNode{Val:1,Next:lA2}

	lB3 :=&ListNode{Val:4,Next:nil}
	lB2 :=&ListNode{Val:3,Next:lB3}
	lB1 :=&ListNode{Val:1,Next:lB2}
	data:=mergeTwoLists(lA1,lB1)

	fmt.Printf("%#v",data)
}