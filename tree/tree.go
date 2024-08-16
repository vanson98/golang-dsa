package tree

import (
	"dsa_demo/queue"
	"fmt"
)

type Node struct {
	data        int
	left, right *Node
}

type Tree struct {
	root     *Node
	capacity int
}

func TestTree() {
	tree := Tree{
		capacity: 10,
	}
	tree.Insert(1)
	tree.Insert(3)
	tree.Insert(33)
	tree.Insert(12)
	tree.Insert(11)
	tree.Insert(87)
	InorderTraversal(tree.root)
}

func (t *Tree) Insert(data int) *Node {
	if t.root == nil {
		t.root = &Node{
			data:  data,
			left:  nil,
			right: nil,
		}
		return t.root
	}
	nodeQueue := queue.InitQueue[Node](t.capacity)
	nodeQueue.Enqueue(*t.root)

	for !nodeQueue.IsEmpty() {
		// check left node
		tempNode := nodeQueue.Dequeue()
		if tempNode.left != nil {
			nodeQueue.Enqueue(*tempNode.left)
		} else {
			tempNode.left = &Node{
				data:  data,
				left:  nil,
				right: nil,
			}
			return tempNode.left
		}

		// check right node
		if tempNode.right != nil {
			nodeQueue.Enqueue(*tempNode.right)
		} else {
			tempNode.right = &Node{
				data:  data,
				left:  nil,
				right: nil,
			}
			return tempNode.right
		}
	}
	fmt.Println("Tree full capacity")
	return nil
}

// Inorder tree traversal (Left - Root - Right)
func InorderTraversal(root *Node) {
	if root == nil {
		return
	}
	InorderTraversal(root.left)
	fmt.Printf("%v ", root.data)
	InorderTraversal(root.right)
}
