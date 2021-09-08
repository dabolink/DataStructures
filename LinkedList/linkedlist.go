package LinkedList

import (
	"errors"
)

var ErrOutOfBounds = errors.New("out of bounds")

type linkedList struct {
	size int
	root *node
}

type node struct {
	next  *node
	value string
}

func (ll linkedList) Length() int {
	return ll.size
}

func (ll *linkedList) Add(value string) {
	newNode := &node{
		next:  nil,
		value: value,
	}

	if ll.root == nil {
		ll.root = newNode
		ll.size += 1
		return
	}

	currentNode := ll.root
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	currentNode.next = newNode
	ll.size += 1
}

func (ll linkedList) Get(index int) (string, error) {
	if index < 0 || index >= ll.Length() {
		return "", ErrOutOfBounds
	}

	i := 0
	currentNode := ll.root
	for currentNode != nil {
		if i == index {
			return currentNode.value, nil
		}
		i++
		currentNode = currentNode.next
	}

	return "", ErrOutOfBounds
}

func (ll *linkedList) Remove(index int) error {
	if index < 0 || index >= ll.Length() {
		return ErrOutOfBounds
	}
	defer func() { ll.size -= 1 }()
	if ll.size == 1 {
		ll.root = nil
		return nil
	}
	currentNode := ll.root
	i := 0
	for currentNode.next.next != nil && i < ll.Length() {
		currentNode = currentNode.next
		i += 1
	}

	currentNode.next = currentNode.next.next
	return nil
}

func New() linkedList {
	return linkedList{
		size: 0,
		root: nil,
	}
}
