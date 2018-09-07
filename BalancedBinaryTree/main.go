package main

import (
	
)
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
 
func isBalanced(root *TreeNode) bool {
	return getDep(root) >= 0
}

func getDep(root *TreeNode) int {
	if root == nil {
		return 0
	}
    
	leftDep := getDep(root.Left)
    rightDep :=getDep(root.Right)
    
	if leftDep < 0 || rightDep < 0 || leftDep-rightDep < -1 || leftDep-rightDep > 1 {
		return -1
	}
	if leftDep > rightDep {
		return leftDep + 1
	}
	return rightDep + 1
}
func main(){

}