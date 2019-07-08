/**
 * 链表中倒数第k个节点
 */
package offer

import (
	"offer100/link"
)

func FindKthToTail(pHead *link.ListNode,k int) *link.ListNode{
	if pHead == nil || k <= 0{
		return nil
	}
	pAhead := pHead
	var pBehind *link.ListNode
	for i:=0; i < k - 1; i++{
		if pAhead != nil{
			pAhead = pAhead.M_pNext
		} else {
			return nil
		}
	}
	pBehind = pHead
	for pAhead.M_pNext != nil{
		pAhead = pAhead.M_pNext
		pBehind = pBehind.M_pNext
	}
	return pBehind


}
