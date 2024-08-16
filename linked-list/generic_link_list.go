package lls

type GNode[T comparable] struct {
	Key       T
	NodeValue interface{}
	Next      *GNode[T]
}

type GLinkList[T comparable] struct {
	Head *GNode[T]
}

func (l *GLinkList[T]) InsertAtTheFront(key T, value interface{}) *GNode[T] {
	newNode := &GNode[T]{Key: key, NodeValue: value, Next: nil}
	if l.Head == nil {
		l.Head = newNode
		return newNode
	}
	newNode.Next = l.Head
	l.Head = newNode
	return newNode
}

func (l *GLinkList[T]) SearchNodeByKey(key T) *GNode[T] {
	currentNode := l.Head
	for currentNode != nil {
		if currentNode.Key == key {
			return currentNode
		} else {
			currentNode = currentNode.Next
		}
	}
	return nil
}
