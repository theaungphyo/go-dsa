package main

import "fmt"

// Node structure
type Node struct {
	value int
	next  *Node
}

// LinkedList structure
type LinkedList struct {
	head *Node
}

// Append Insert at end
func (list *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Display the list
func (list *LinkedList) Display() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.Append(10)
	list.Append(20)
	list.Append(30)
	list.Display() // 10 -> 20 -> 30 -> nil
}
