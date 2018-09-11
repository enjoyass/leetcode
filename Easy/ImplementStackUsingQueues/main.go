package main

import (
	"fmt"
)


type MyStack struct {
    Queues []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.Queues=append(this.Queues,x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmp:= this.Queues[len(this.Queues)-1]
	this.Queues=this.Queues[:len(this.Queues)-1]
	return tmp
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Queues[len(this.Queues)-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    if len(this.Queues)==0 {
		return true
	}
	return false
}

func main(){
	stack :=Constructor()
	stack.Push(1)
	stack.Push(2); 
	fmt.Println(stack.Top())  // returns 2
	fmt.Println(stack.Pop())   // returns 2
	fmt.Println(stack.Empty()) // returns false
}