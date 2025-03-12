package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	// set up a dummy node here
	dummy := &ListNode{}
	curr := dummy
	carry := 0

	// iterate through both list and the carry
	for list1 != nil || list2 != nil || carry != 0 {
		sum := carry

		if list1 != nil{
			sum += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			sum += list2.Val
			list2 = list2.Next
		}

		carry = sum /10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	return dummy.Next
}

// helper function to print a LinkedList
func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		if node.Next != nil {
			fmt.Println("->")
		}
		node = node.Next
	}
}

func main() {
    x := &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
    y := &ListNode{Val: 5, Next: &ListNode{Val: 0, Next: &ListNode{Val: 2}}}
    sum := addTwoNumbers(x, y)
	printList(sum)
}
