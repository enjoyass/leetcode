package main
import (
	"fmt"
)
type MyQueue struct {
    Queues []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.Queues=append(this.Queues,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	tmp:=this.Queues[0]
	this.Queues=this.Queues[1:]
	return tmp
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.Queues[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    if len(this.Queues)==0 {
		return true
	}
	return false
}

func main(){
	stack :=Constructor()
	stack.Push(1)
	stack.Push(2); 
	fmt.Println(stack.Peek())  // returns 2
	fmt.Println(stack.Pop())   // returns 2
	fmt.Println(stack.Empty()) // returns false
}