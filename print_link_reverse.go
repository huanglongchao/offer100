package main

import (
	"fmt"
	"offer100/link"
)

func main(){
	var pHead *link.ListNode
	link.AddToTail(&pHead,1)
	link.AddToTail(&pHead,2)
	link.AddToTail(&pHead,3)
	link.AddToTail(&pHead,4)
	link.AddToTail(&pHead,5)
	link.PrintListReversingly_Iteratively(pHead)
	fmt.Printf("\n")
	link.RemoveNode(&pHead,3)
	link.PrintListReversingly_Iteratively(pHead)
}
