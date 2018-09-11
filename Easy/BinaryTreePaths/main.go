package main

import (
	"strings"
	"strconv"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func main(){

}
func binaryTreePaths(root *TreeNode) []string {
    result := &[]string{}
    path := []int{}
    binaryTreePathsHelper(root, path, result)
    return *result
}
func binaryTreePathsHelper(root *TreeNode, path []int, result *[]string) {
    if root == nil {
        return 
    }
    path = append(path, root.Val)
    if root.Left == nil && root.Right == nil {
        if len(path) > 0 {
            s := []string{}
            for i := 0; i < len(path); i++ {
                s = append(s, strconv.Itoa(path[i]))
            }
            *result = append(*result, strings.Join(s, "->"))
        }
        return 
    }
    binaryTreePathsHelper(root.Left, path, result)
    binaryTreePathsHelper(root.Right, path, result)
} 
 
