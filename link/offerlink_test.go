package link

import (
	"fmt"
	"testing"
)

func TestAddToTail(t *testing.T) {
	var pHead *ListNode
	AddToTail(&pHead,1)
	AddToTail(&pHead,2)
	AddToTail(&pHead,3)
	AddToTail(&pHead,4)
	AddToTail(&pHead,5)
	PrintListReversingly_Iteratively(pHead)
}

func TestPrintListReversingly_Iteratively(t *testing.T) {
	pHead := &ListNode{
		M_nValue: 1,
	}
	PrintListReversingly_Iteratively(pHead)
}

func TestRemoveNode(t *testing.T) {
	var pHead *ListNode
	AddToTail(&pHead,1)
	AddToTail(&pHead,2)
	AddToTail(&pHead,3)
	AddToTail(&pHead,4)
	AddToTail(&pHead,5)
	PrintListReversingly_Iteratively(pHead)
	fmt.Printf("\n")
	RemoveNode(&pHead,3)
	PrintListReversingly_Iteratively(pHead)
}