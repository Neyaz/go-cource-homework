package doubly_linked_list

import (
	"errors"
)

// Node of linked list
type Node struct {
	value      string
	next, prev *Node
}

// DoublyLinkedList doubly linked list
type DoublyLinkedList struct {
	head, tail *Node
}

// First gets first element of list
func (l *DoublyLinkedList) First() *Node {
	return l.head
}

// Last gets last element of list
func (l *DoublyLinkedList) Last() *Node {
	return l.tail
}

// Next gets next node
func (n *Node) Next() *Node {
	return n.next
}

// Prev gets previous node
func (n *Node) Prev() *Node {
	return n.prev
}

// HeadPush add new node to the head of the list
func (l *DoublyLinkedList) HeadPush(v string) *DoublyLinkedList {
	n := &Node{value: v}
	if l.tail == nil {
		l.tail = n
	}

	if l.head != nil {
		if l.head.next == nil {
			l.tail = l.head
		}

		l.head.prev = n
		n.next = l.head
	}

	l.head = n
	return l
}

// TailPush add new node to the end of the list
func (l *DoublyLinkedList) TailPush(v string) *DoublyLinkedList {
	n := &Node{value: v}
	if l.head == nil && l.tail == nil {
		l.head = n
	} else if l.tail == nil {
		if l.tail.prev == nil {
			l.head = l.tail
		}
		n.prev = l.head
		l.head.next = n
	} else {
		n.prev = l.tail
		l.tail.next = n
	}
	l.tail = n
	return l
}

// Find node in the list
func (l *DoublyLinkedList) Find(value string) *Node {
	for n := l.First(); n != nil; n = n.Next() {
		if n.value == value {
			return n
		}
	}
	return nil
}

// HeadPop removes head
func (l *DoublyLinkedList) HeadPop() (*DoublyLinkedList, error) {
	if l.head == nil && l.tail == nil {
		return nil, errors.New("Empty list")
	}

	l.head.next.prev = nil
	l.head = l.head.next
	l.head.prev = nil
	return l, nil
}

// TailPop removes head
func (l *DoublyLinkedList) TailPop() (*DoublyLinkedList, error) {
	if l.head == nil && l.tail == nil {
		return nil, errors.New("Empty list")
	}

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.prev.next = nil
		l.tail = l.tail.prev
	}
	return l, nil
}

// Delete node from the list
func (l *DoublyLinkedList) Delete(value string) *DoublyLinkedList {
	node := l.Find(value)
	if node != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	return l
}
