package link

import "fmt"

type ListNode struct {
	M_nValue int
	M_pNext  *ListNode
}

func AddToTail(pHead **ListNode, value int){

	pNew := &ListNode{}
	pNew.M_nValue = value
	pNew.M_pNext = nil

	if *pHead == nil{

		*pHead = pNew
	} else {
		pNode := *pHead
		for pNode.M_pNext != nil  {
			pNode = pNode.M_pNext
		}
		pNode.M_pNext = pNew
	}
}

func PrintListReversingly_Iteratively(pHead *ListNode){
	if pHead != nil{
		if pHead.M_pNext != nil{
			PrintListReversingly_Iteratively(pHead.M_pNext)
		}
		fmt.Printf("%v \t",pHead.M_nValue)
	}
}

func RemoveNode(pHead **ListNode, value int){
	if pHead == nil || *pHead == nil{
		return
	}
	var pToBeDeleted *ListNode
	if (*pHead).M_nValue == value{
		pToBeDeleted = *pHead
		*pHead = (*pHead).M_pNext
	} else {
		pNode := *pHead
		for pNode.M_pNext != nil && pNode.M_pNext.M_nValue != value{
			pNode = pNode.M_pNext
		}
		if pNode.M_pNext != nil && pNode.M_pNext.M_nValue == value{
			pToBeDeleted = pNode.M_pNext
			pNode.M_pNext = pNode.M_pNext.M_pNext
		}
	}
	if pToBeDeleted != nil{
		pToBeDeleted = nil
	}
}

/**
 * 给定单向链表头指针和一个节点指针，o(1)删除链表节点
 */
func deleteLinkPNode(pHead **ListNode, pNode *ListNode){
	if pHead == nil || pNode == nil{
		return
	}
	if pNode.M_pNext != nil{
		pNode.M_nValue = pNode.M_pNext.M_nValue
		pNode.M_pNext = pNode.M_pNext.M_pNext
	} else if *pHead == pNode{
		pNode = nil
		*pHead = nil
	} else {
		pTmp := *pHead
		for pTmp.M_pNext != pNode{
			pTmp = pTmp.M_pNext
		}
		pTmp.M_pNext = nil
	}
}

/**
 * 在一个排序的链表中，删除重复的节点
 */

func deleteDupNode(pHead **ListNode){
	if pHead == nil || *pHead == nil {
		return
	}
	var preNode *ListNode
	pNode := *pHead
	for pNode != nil{
		pNext := pNode.M_pNext
		needDelete := false
		if pNext != nil && pNext.M_nValue == pNode.M_nValue {
			needDelete = true
		}
		if !needDelete{
			preNode = pNode
			pNode = pNode.M_pNext
		} else {
			value := pNode.M_nValue
			pToBeDel := pNode
			for pToBeDel != nil && pToBeDel.M_nValue == value  {
				pNext = pToBeDel.M_pNext
				pToBeDel = pNext
			}

			if preNode == nil{
				*pHead = pNext
			} else {
				preNode.M_pNext = pNext
			}
			pNode = pNext
		}
	}
}