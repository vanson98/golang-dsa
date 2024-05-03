package lls

type CircularLinkList struct {
	last *Node
}

var clLinkList CircularLinkList

func TestCircularLinkList() {

	clLinkList.insertAtTheFront("Nguyen")
	clLinkList.insertAtTheFront("Van")
	clLinkList.insertAtTheFront("Son")
	scanCircleLinkList(clLinkList.last)
}

func (l *CircularLinkList) insertAtTheFront(value string) *Node {
	newNode := &Node{data: value, next: nil}
	if l.last == nil {
		l.last = newNode
		newNode.next = newNode
		return newNode
	}
	newNode.next = l.last.next
	l.last.next = newNode
	return newNode
}

func scanCircleLinkList(node *Node) {
	currentNode := node
	for {
		PrintNodeInfo(currentNode)
		currentNode = currentNode.next
		if currentNode == node {
			return
		}
	}
}
