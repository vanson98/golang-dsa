package lls

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	data string
	next *Node
}

type LinkList struct {
	head *Node
}

var linkList LinkList

func Test() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter new node data: ")
	// text, _ := reader.ReadString('\n')
	// text = strings.Replace(text, "\r\n", "", -1)
	// Add new node
	linkList.InsertAtTheFront("Nguyen")
	vanNode := linkList.InsertAtTheFront("Van")
	linkList.InsertAtTheFront("Son")
	linkList.InsertAtTheMiddle(vanNode, "123")
	linkList.InsertAtTheEnd("DT")
	ScanLinkList(linkList.head)

	// delete node
	linkList.DeleteAtTheFront()
	fmt.Println("After Delete Head")
	ScanLinkList(linkList.head)

	// search node
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter search value: ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	searchedNode := linkList.SearchByValue(text)
	PrintNodeInfo(searchedNode)

	// count node
	fmt.Printf("Total node of link list is: %d\n", linkList.GetLength())

	// delete at the middle
	fmt.Println("Delete middle node")
	linkList.DeleteAtTheMiddle(2)
	ScanLinkList(linkList.head)
}

func (l *LinkList) InsertAtTheFront(value string) *Node {
	newNode := &Node{data: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return newNode
	}
	newNode.next = l.head
	l.head = newNode
	return newNode
}

func (l *LinkList) InsertAtTheMiddle(prev_node *Node, value string) {
	newNode := &Node{data: value, next: prev_node.next}
	prev_node.next = newNode
}

func (l *LinkList) InsertAtTheEnd(value string) {
	newNode := &Node{data: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	currentNode := l.head
	for {
		if currentNode.next == nil {
			currentNode.next = newNode
			break
		}
		currentNode = currentNode.next
	}
}

func (l *LinkList) DeleteAtTheFront() {
	currentHead := l.head
	l.head = currentHead.next
	currentHead = nil
}

func (l *LinkList) DeleteAtTheMiddle(nodeIndex int) {
	currentNode := l.head
	for i := 1; i < nodeIndex && currentNode.next != nil; i++ {
		currentNode = currentNode.next
	}
	temp := currentNode.next
	currentNode.next = temp.next
	temp = nil
}

func (l *LinkList) GetLength() int {
	currentNode := l.head
	count := 0
	if currentNode == nil {
		return count
	}
	for {
		count++
		if currentNode.next == nil {
			break
		}
		currentNode = currentNode.next
	}
	return count
}

func (l *LinkList) SearchByValue(value string) *Node {
	currentNode := l.head
	for {
		if currentNode.data == value {
			return currentNode
		}
		currentNode = currentNode.next
	}
}

func ScanLinkList(node *Node) {
	PrintNodeInfo(node)
	if node.next != nil {
		ScanLinkList(node.next)
	}
	return
}

func PrintNodeInfo(node *Node) {
	fmt.Printf("type:%T     &pointer_of_pointer:%p      pointer_address:%p      pointer_value:%v\n", node, &node, node, *node)
}
