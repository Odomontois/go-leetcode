package main

import "fmt"

type X struct{}

func main() {
	var node *ListNode
	node = node.Next
	fmt.Println("Hello")
}

type ListNode struct {
	Val  int
	Next *ListNode
}
