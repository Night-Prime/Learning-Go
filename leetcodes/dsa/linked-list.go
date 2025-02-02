//Learning about LinkedList today.
package main

import "fmt"

// initialize the node & linkedlist

type Node struct{
	data int
	next *Node
}

type LinkedList struct{
	head *Node
}

// Insertion at the start

func (list *LinkedList) insertAtStart (data int) {
	fmt.Println("Insertion at the beginning \n")
	// if the head is empty, set the new node to it
	if list.head == nil {
		newNode := &Node{data:data, next:nil}
		list.head = newNode
		return
	}
	// else append the incoming value to the head
	newNode := &Node{data:data, next:list.head}
	list.head = newNode
}


// Insertion at the end

func (list *LinkedList) insertAtEnd (data int){
	fmt.Println("Insertion at the end \n")
	// create a new node value first
	newNode := &Node{data:data, next:nil}

	// if the list is empty, set the new node to it
	if list.head == nil {
		list.head = newNode
		return
	}

	// set a current value to the head
	currentVal := list.head

	// iterating through the whole list to find the last node
	for currentVal.next != nil {
		currentVal = currentVal.next
	}

	// then set inserted value into the last node
	currentVal.next = newNode;
}


// Print Linked List

func (list *LinkedList) printLinkedList() {
	fmt.Println("Printing the linked list \n")

	// set current pointer to the head
	currentVal := list.head
	// iterate through list and print current value
	for currentVal != nil{
		fmt.Printf("%d -> \n", currentVal.data)
		currentVal = currentVal.next 
	}
	fmt.Println()
}


// Deletion at the start

func (list *LinkedList) deleteAtStart() {
	fmt.Println("Deletion at the start\n")
	// if the head is empty
	if list.head == nil {
		fmt.Println("Empty Linked List")
		return
	}
	// else simply set the next value as the head
	list.head = list.head.next
}

// Deletion at the end 
func (list *LinkedList) deleteAtEnd() {
	fmt.Println("Deletion at the end \n")
	// if the head is empty
	if list.head == nil{
		fmt.Println("Empty Linked List")
		return
	}
	// if the it's only a value in the linkedList
	if list.head.next == nil{
		list.head = nil
		fmt.Println("Only one Value in this Linked List")
		return
	}

	var currentVal *Node = list.head
	// check for the second to the last node
	for currentVal.next.next != nil {
		currentVal = currentVal.next
	}
	currentVal.next = nil
}


// Print the count

func(list *LinkedList) countNodes() (count int) {
	currentVal := list.head
	for currentVal != nil {
		currentVal = currentVal.next
		count++
	}
	fmt.Printf("The Node Count is %d \n", count)
	return
}


// Finding at value at index
func(list *LinkedList) findNodeAtIndex(index int) *Node {
	fmt.Printf("Searching for Node value at %d \n", index)
	count := 1
	currentVal := list.head
	for currentVal.next != nil {
		count++
		currentVal = currentVal.next
	}

	if index <= 0 || index > count {
		fmt.Printf("The index is out of range")
		return nil
	}

	currentVal = list.head
	for count = 1; count < index; count++ {
		currentVal = currentVal.next
	}
	fmt.Printf("The Node at %d is %d \n", index, currentVal.data)
	return currentVal
}


func main(){
	fmt.Println("Setting up Linked List \n")

	fmt.Println("################################################ \n")

	sampleLinkedList := LinkedList{}

	sampleLinkedList.insertAtStart(2)
	sampleLinkedList.insertAtEnd(4)
	sampleLinkedList.insertAtEnd(6)
	sampleLinkedList.insertAtEnd(8)
	sampleLinkedList.insertAtEnd(10)
	sampleLinkedList.insertAtEnd(6)
	sampleLinkedList.insertAtEnd(10)

	sampleLinkedList.deleteAtStart()
	sampleLinkedList.deleteAtEnd()

	sampleLinkedList.countNodes()
	sampleLinkedList.findNodeAtIndex(2)

	sampleLinkedList.printLinkedList()

}