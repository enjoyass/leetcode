package main
type ListNode struct {
    Val int
    Next *ListNode
}
func main(){

}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}