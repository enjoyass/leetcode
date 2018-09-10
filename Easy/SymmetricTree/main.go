package main
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
 }
func isSymmetric(root *TreeNode) bool {
    if root ==nil {
		return true
	}
    return doSymmetric(root.Left,root.Right)
}

func doSymmetric(leftNode *TreeNode,rightNode *TreeNode)bool{
	if leftNode ==nil && rightNode==nil {
		return true
	}
	if leftNode == nil {
		return false
	}
	if rightNode == nil {
		return false
	}
	
	return (leftNode.Val == rightNode.Val) && doSymmetric(leftNode.Left,rightNode.Right) &&doSymmetric(leftNode.Right,rightNode.Left)
}

func main(){

}