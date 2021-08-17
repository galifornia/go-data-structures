package main

import "fmt"

type TNode struct {
	key   int
	left  *TNode
	right *TNode
}

type BinaryTree struct {
	root *TNode
}

// Insert: receives a value and search for the right place to insert it down the tree.
// If v is larger goes to the right TNode, otherwise goes to left TNode
func (tree *BinaryTree) Insert(v int) {
	p := tree.root

	for {
		if p.key < v {
			if p.right == nil {
				p.right = &TNode{key: v}
				return
			}
			p = p.right
		} else if p.key >= v {
			if p.left == nil {
				p.left = &TNode{key: v}
				return
			}
			p = p.left
		}
	}
}

// print: utility function that prints all keys present in the tree.
func Print(node *TNode) {
	if node == nil {
		return
	}

	fmt.Println(node.key)
	if node.left != nil {
		print(node.left)
	}
	if node.right != nil {
		print(node.right)
	}
}

// Search: gets k input and returns true if there is key with that value
func (tree *BinaryTree) Search(k int) (bool, int) {
	p := tree.root
	count := 0

	for p != nil {
		count++
		if p.key == k {
			return true, count
		}

		if p.key < k {
			p = p.right
		} else if p.key > k {
			p = p.left
		}
	}

	return false, count
}

// func main() {
// 	tree := BinaryTree{}
// 	tree.root = &TNode{key: 23}

// 	nodeValues := []int{100, 21, 3, 17, 9, 204, 150, 53, 636, 664, 1, 45}
// 	for _, v := range nodeValues {
// 		tree.Insert(v)
// 	}
// 	fmt.Println(tree.root)
// 	print(tree.root)

// 	fmt.Println(tree.Search(27))
// 	fmt.Println(tree.Search(100))
// 	fmt.Println(tree.Search(10))
// 	fmt.Println(tree.Search(3))
// 	fmt.Println(tree.Search(204))
// 	fmt.Println(tree.Search(13))
// 	fmt.Println(tree.Search(636))
// 	fmt.Println(tree.Search(1))
// }
