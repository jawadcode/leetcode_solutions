package main

import (
	"fmt"
	"math"
)

// ListNode is a linked list node, doofus
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func arrToLinked(arr []int) *ListNode {
	arr = reverse(arr)
	node := &ListNode{
		Val:  arr[0],
		Next: nil,
	}
	for _, v := range arr[1:] {
		node = &ListNode{
			Val:  v,
			Next: node,
		}
	}
	return node
}

func linkedToInt(l *ListNode) uint64 {
	var res uint64
	var arr []int
	for temp := l; temp != nil; temp = temp.Next {
		arr = append(arr, temp.Val)
	}
	for i, v := range arr {
		res += uint64(uint64(v) * uint64(math.Pow10(i)))
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, temp *ListNode
	temp = &ListNode{0, &ListNode{}}
	res = temp
	totalSum := 0
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		totalSum = carry
		if l1 != nil {
			totalSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			totalSum += l2.Val
			l2 = l2.Next
		}
		temp.Next = &ListNode{Val: totalSum % 10}
		carry = totalSum / 10
		temp = temp.Next
	}
	return res.Next
}

// This function is just for demonstratory purposes and would of course fail with large enough numbers
func main() {
	l1 := arrToLinked([]int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})
	l2 := arrToLinked([]int{5, 6, 4})
	l3 := addTwoNumbers(l1, l2)
	fmt.Println("(", linkedToInt(l1), "+", linkedToInt(l2), "=", linkedToInt(l3), ")")
	fmt.Println(*l3)
}
