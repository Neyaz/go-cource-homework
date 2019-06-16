package doubly_linked_list

import (
	"errors"
)

// Node of linked list
type Node struct {
  value string
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
	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
		n.prev = l.tail
	}
	l.tail = n
	return l
}

// TailPush add new node to the end of the list
func (l *DoublyLinkedList) TailPush(v string) (*DoublyLinkedList, error) {
  n := &Node{value: v}
  if l.head == nil && l.tail == nil {
      return nil, errors.New("No head in the list")
  }

  if l.tail == nil {
    n.prev = l.head
    l.head.next = n
  } else {
    n.prev = l.tail
    l.tail.next = n
  }
	l.tail = n
	return l, nil
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
  if l.head == nil {
      return nil, errors.New("No head in the list")
  }

  l.head.next.prev = nil
  l.head = l.head.next
	return l, nil
}

// TailPop removes head
func (l *DoublyLinkedList) TailPop() (*DoublyLinkedList, error) {
  if l.head == nil {
      return nil, errors.New("No head in the list")
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