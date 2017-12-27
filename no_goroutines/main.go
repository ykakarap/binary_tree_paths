package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	a := &TreeNode{
		Val: 5,
	}
	b := &TreeNode{
		Val: 3,
	}
	c := &TreeNode{
		Val:   2,
		Right: a,
	}
	d := &TreeNode{
		Val:   1,
		Left:  c,
		Right: b,
	}
	fmt.Println("Paths are")
	fmt.Println(binaryTreePaths(nil))
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return make([]string, 0)
	} else {
		return getPaths(root, "", make([]string, 0))
	}
}

func getPaths(node *TreeNode, prefix string, paths []string) []string {
	if prefix == "" {
		prefix = strconv.Itoa(node.Val)
	} else {
		prefix = prefix + "->" + strconv.Itoa(node.Val)
	}
	var lPaths []string
	var rPaths []string
	if node.Left == nil && node.Right == nil {
		// LEeaf
		paths = append(paths, prefix)
	}
	if node.Left != nil {
		lPaths = getPaths(node.Left, prefix, paths)
	}
	if node.Right != nil {
		rPaths = getPaths(node.Right, prefix, paths)
	}

	sPaths := append(lPaths, rPaths...)
	tPaths := append(paths, sPaths...)

	return tPaths
}
