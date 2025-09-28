package trees

import (
	"fmt"
	"strings"
)

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree represents a binary tree data structure
type BinaryTree struct {
	Root *TreeNode
}

// NewBinaryTree creates a new empty binary tree
func NewBinaryTree() *BinaryTree {
	return &BinaryTree{Root: nil}
}

// Insert inserts a value into the binary search tree
func (bt *BinaryTree) Insert(value int) {
	bt.Root = bt.insertRec(bt.Root, value)
}

func (bt *BinaryTree) insertRec(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}

	if value < node.Value {
		node.Left = bt.insertRec(node.Left, value)
	} else if value > node.Value {
		node.Right = bt.insertRec(node.Right, value)
	}

	return node
}

// Search searches for a value in the binary search tree
func (bt *BinaryTree) Search(value int) bool {
	return bt.searchRec(bt.Root, value)
}

func (bt *BinaryTree) searchRec(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}

	if value == node.Value {
		return true
	}

	if value < node.Value {
		return bt.searchRec(node.Left, value)
	}

	return bt.searchRec(node.Right, value)
}

// Delete removes a value from the binary search tree
func (bt *BinaryTree) Delete(value int) {
	bt.Root = bt.deleteRec(bt.Root, value)
}

func (bt *BinaryTree) deleteRec(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return node
	}

	if value < node.Value {
		node.Left = bt.deleteRec(node.Left, value)
	} else if value > node.Value {
		node.Right = bt.deleteRec(node.Right, value)
	} else {
		// Node to be deleted found
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// Node with two children: get inorder successor
		node.Value = bt.minValue(node.Right)
		node.Right = bt.deleteRec(node.Right, node.Value)
	}

	return node
}

func (bt *BinaryTree) minValue(node *TreeNode) int {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

// InorderTraversal returns values in inorder traversal (left, root, right)
func (bt *BinaryTree) InorderTraversal() []int {
	var result []int
	bt.inorderRec(bt.Root, &result)
	return result
}

func (bt *BinaryTree) inorderRec(node *TreeNode, result *[]int) {
	if node != nil {
		bt.inorderRec(node.Left, result)
		*result = append(*result, node.Value)
		bt.inorderRec(node.Right, result)
	}
}

// PreorderTraversal returns values in preorder traversal (root, left, right)
func (bt *BinaryTree) PreorderTraversal() []int {
	var result []int
	bt.preorderRec(bt.Root, &result)
	return result
}

func (bt *BinaryTree) preorderRec(node *TreeNode, result *[]int) {
	if node != nil {
		*result = append(*result, node.Value)
		bt.preorderRec(node.Left, result)
		bt.preorderRec(node.Right, result)
	}
}

// PostorderTraversal returns values in postorder traversal (left, right, root)
func (bt *BinaryTree) PostorderTraversal() []int {
	var result []int
	bt.postorderRec(bt.Root, &result)
	return result
}

func (bt *BinaryTree) postorderRec(node *TreeNode, result *[]int) {
	if node != nil {
		bt.postorderRec(node.Left, result)
		bt.postorderRec(node.Right, result)
		*result = append(*result, node.Value)
	}
}

// Height returns the height of the tree
func (bt *BinaryTree) Height() int {
	return bt.heightRec(bt.Root)
}

func (bt *BinaryTree) heightRec(node *TreeNode) int {
	if node == nil {
		return -1
	}

	leftHeight := bt.heightRec(node.Left)
	rightHeight := bt.heightRec(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// IsEmpty checks if the tree is empty
func (bt *BinaryTree) IsEmpty() bool {
	return bt.Root == nil
}

// Size returns the number of nodes in the tree
func (bt *BinaryTree) Size() int {
	return bt.sizeRec(bt.Root)
}

func (bt *BinaryTree) sizeRec(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + bt.sizeRec(node.Left) + bt.sizeRec(node.Right)
}

// String returns a string representation of the tree
func (bt *BinaryTree) String() string {
	if bt.Root == nil {
		return "Empty tree"
	}

	var result strings.Builder
	result.WriteString("Binary Tree:\n")
	bt.printTree(bt.Root, "", true, &result)
	return result.String()
}

func (bt *BinaryTree) printTree(node *TreeNode, prefix string, isLast bool, result *strings.Builder) {
	if node != nil {
		result.WriteString(prefix)
		if isLast {
			result.WriteString("└── ")
		} else {
			result.WriteString("├── ")
		}
		result.WriteString(fmt.Sprintf("%d\n", node.Value))

		children := []*TreeNode{node.Left, node.Right}
		for i, child := range children {
			if child != nil {
				extension := "│   "
				if isLast {
					extension = "    "
				}
				bt.printTree(child, prefix+extension, i == len(children)-1, result)
			}
		}
	}
}
