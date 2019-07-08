/**
 * 链表中 环的入口节点
 */
package offer

import "offer100/link"

func MeetingNode(pHead *link.ListNode) *link.ListNode{
	if pHead == nil{
		return nil
	}
	pSlow := pHead.M_pNext
	if pSlow == nil{
		return nil
	}
	pFast := pSlow.M_pNext
	for pFast != nil && pSlow != nil{
		if pFast == pSlow{
			return pFast
		}
		pSlow = pSlow.M_pNext
		pFast = pFast.M_pNext
		if pFast != nil{
			pFast = pFast.M_pNext
		}
	}
	return nil
}

func EntryNodeOfLoop(pHead *link.ListNode) *link.ListNode{
	meetingNode := MeetingNode(pHead);
	if meetingNode == nil{
		return nil
	}
	//得到环中节点的数目
	nodesInLoop := 1
	pNode1 := meetingNode
	for pNode1.M_pNext != meetingNode{
		pNode1 = pNode1.M_pNext
		nodesInLoop++
	}
	//先移动pNode1, 次数为环中节点的数目
	pNode1 = pHead
	for i:=0; i < nodesInLoop; i++{
		pNode1 = pNode1.M_pNext
	}
	//再移动pNode1 和 pNode2
	pNode2 := pHead
	for pNode1 != pNode2{
		pNode1 = pNode1.M_pNext
		pNode2 = pNode2.M_pNext
	}
	return pNode1
}