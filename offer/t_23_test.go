package offer

import (
	"fmt"
	"offer100/link"
	"testing"
)

func TestEntryNodeOfLoop(t *testing.T) {
	var pHead *link.ListNode
	link.AddToTail(&pHead,1)
	link.AddToTail(&pHead,2)
	link.AddToTail(&pHead,3)
	link.AddToTail(&pHead,4)
	link.AddToTail(&pHead,5)
	link.PrintListReversingly_Iteratively(pHead)
	pHead.M_pNext.M_pNext.M_pNext = pHead.M_pNext.M_pNext
	fmt.Println(EntryNodeOfLoop(pHead).M_nValue)
}