package data_structures

import "fmt"

type Treenode struct {
	data  int
	left  *Treenode
	right *Treenode
}

type BinarySearchTree struct {
	root *Treenode
}

func (treenode_ *Treenode) TestTree() {

}
func CallTreeNode() {
	var treenode *Treenode = &Treenode{}
	treenode.TestTree()
}

// func (bst *BinarySearchTree) NewBST() Treenode {
// 	return Treenode{root: }
// }

func NewRoot(bst *BinarySearchTree) *Treenode {
	return bst.root
}

// func NewTreeNode() *Treenode {
// 	var newnode *Treenode = &Treenode{}
// 	return newnode
// }

func newTreeNode(key int) *Treenode {
	var node *Treenode = &Treenode{}
	node.data = key
	node.left = node.right
	node.right = nil
	return node
}

func InsertTree(tnode *Treenode, key int) *Treenode {
	if tnode == nil {
		return newTreeNode(key)
	}
	if key < tnode.data {
		tnode.left = InsertTree(tnode.left, key)
	} else if key > tnode.data {
		tnode.right = InsertTree(tnode.right, key)
	}
	return tnode
}

func (bst *BinarySearchTree) Insert(key int) {
	bst.insertNode(bst.root, key)
}

func (bst *BinarySearchTree) insertNode(tnode *Treenode, key int) *Treenode {
	if bst.root == nil {
		return newTreeNode(key)
	}
	if tnode == nil {
		return newTreeNode(key)
	}
	if key < tnode.data {
		tnode.left = bst.insertNode(tnode.left, key)
	} else if key > tnode.data {
		tnode.right = bst.insertNode(tnode.right, key)
	}
	return tnode
}

func InorderTree(tnode *Treenode) {
	if tnode != nil {
		InorderTree(tnode.left)
		fmt.Println(tnode.data)
		InorderTree(tnode.right)
	}
}

func (bst *BinarySearchTree) InorderTree(tnode *Treenode) {
	if bst.root != nil {
		bst.InorderTree(bst.root.left)
		fmt.Println(bst.root.data)
		bst.InorderTree(tnode.right)
	}
}
