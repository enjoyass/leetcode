package main

import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
 func bfs(s *[][]int, level int, root *TreeNode)  {
	if root == nil {
		return
	}
	if len(*s) == level {
		*s = append(*s, []int{})
	}
	(*s)[level] = append((*s)[level], root.Val)
	for _, v := range []*TreeNode{root.Left, root.Right} {
		bfs(s, level + 1, v)
	}
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var s [][]int
	bfs(&s, 0, root)
	for i := 0; i < len(s) / 2; i++ {
		s[i], s[len(s) - i - 1] = s[len(s) - i - 1], s[i]
	}
	return s
}

func main(){
	lfl:=&TreeNode{Val:3,Left:nil,Right:nil}
	lfr:=&TreeNode{Val:3,Left:nil,Right:nil}
	ril:=&TreeNode{Val:3,Left:nil,Right:nil}
	rir:=&TreeNode{Val:3,Left:nil,Right:nil}
	lef:=&TreeNode{Val:2,Left:lfl,Right:lfr}
	right:=&TreeNode{Val:2,Left:ril,Right:rir}
	head := &TreeNode{Val:1,Left:lef,Right:right}
	fmt.Println(levelOrderBottom(head))
}