package main

import (
	"fmt"
)
 // Definition for singly-linked list.
 type ListNode struct {
     Val int
     Next *ListNode
 }
 
func main() {
	l5:=&ListNode{Val:3,Next:nil}
	l4:=&ListNode{Val:3,Next:l5}
	l3:=&ListNode{Val:2,Next:l4}
	l2:=&ListNode{Val:1,Next:l3}
	l1:=&ListNode{Val:1,Next:l2}

	data := deleteDuplicates(l1)

	for data != nil {
		fmt.Println(data.Val)
		data=data.Next
	}
}
//在原有的基础上修改，速度快
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		next := cur.Next
		if next == nil {
			return head
		}
		for cur.Val == next.Val {
			next = next.Next
			if next == nil {
				break
			}
		}
		cur.Next = next
		cur = cur.Next
	}
	return head
}


//速度慢，去除去重，重新建链表
func deleteDuplicates2(head *ListNode) *ListNode {
	newMap :=make(map[int]int)
	var num []int
	for head != nil {
		if _,ok:=newMap[head.Val];!ok {
			newMap[head.Val]=head.Val
			num=append(num,head.Val)
		}
		head=head.Next
	}
	tmp:=head
	for i:=0;i<len(num);i++{
		node := &ListNode{Val:num[i],Next:nil}
		if head == nil {
			head=node
		}else{
			tmp.Next=node
		}
		tmp=node
	}
    return head
}