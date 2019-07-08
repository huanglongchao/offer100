package offer

import (
	"fmt"
	"offer100/link"
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	var pHead *link.ListNode
	link.AddToTail(&pHead,1)
	link.AddToTail(&pHead,2)
	link.AddToTail(&pHead,3)
	link.AddToTail(&pHead,4)
	link.AddToTail(&pHead,5)
	link.PrintListReversingly_Iteratively(pHead)
	pNode := FindKthToTail(pHead,3)
	fmt.Println()
	link.PrintListReversingly_Iteratively(pNode)
}
