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
	fmt.Println(binaryTreePaths(d))

}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return make([]string, 0)
	}

	ch := make(chan []string)
	go getPathstoLeefs(root, "", ch)
	return <-ch
}

func getPathstoLeefs(node *TreeNode, basepath string, ch chan []string) {
	var path string
	if basepath == "" {
		path = strconv.Itoa(node.Val)
	} else {
		path = basepath + "->" + strconv.Itoa(node.Val)
	}

	if node.Left == nil && node.Right == nil {
		// this node is a leef
		ch <- []string{
			path,
		}
	} else {
		lch := make(chan []string)
		rch := make(chan []string)
		if node.Left != nil {
			// Get paths from left side
			go getPathstoLeefs(node.Left, path, lch)
		} else {
			go func(l chan []string) {
				l <- nil
			}(lch)
		}

		if node.Right != nil {
			// get paths from right side
			go getPathstoLeefs(node.Right, path, rch)
		} else {
			go func(r chan []string) {
				r <- nil
			}(rch)
		}

		// return the merge of both sides
		// lPaths, rPaths := <-lch, <-rch
		lPaths := <-lch
		rPaths := <-rch

		ch <- append(lPaths, rPaths...)
	}
}
