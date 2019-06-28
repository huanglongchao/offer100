package link

import "fmt"

type ListNode struct {
	m_nValue int
	m_pNext *ListNode
}

func AddToTail(pHead **ListNode, value int){

	pNew := &ListNode{}
	pNew.m_nValue = value
	pNew.m_pNext = nil

	if *pHead == nil{

		*pHead = pNew
	} else {
		pNode := *pHead
		for pNode.m_pNext != nil  {
			pNode = pNode.m_pNext
		}
		pNode.m_pNext = pNew
	}
}

func PrintListReversingly_Iteratively(pHead *ListNode){
	if pHead != nil{
		if pHead.m_pNext != nil{
			PrintListReversingly_Iteratively(pHead.m_pNext)
		}
		fmt.Printf("%v \t",pHead.m_nValue)
	}
}

func RemoveNode(pHead **ListNode, value int){
	if pHead == nil || *pHead == nil{
		return
	}
	var pToBeDeleted *ListNode
	if (*pHead).m_nValue == value{
		pToBeDeleted = *pHead
		*pHead = (*pHead).m_pNext
	} else {
		pNode := *pHead
		for pNode.m_pNext != nil && pNode.m_pNext.m_nValue != value{
			pNode = pNode.m_pNext
		}
		if pNode.m_pNext != nil && pNode.m_pNext.m_nValue == value{
			pToBeDeleted = pNode.m_pNext
			pNode.m_pNext = pNode.m_pNext.m_pNext
		}
	}
	if pToBeDeleted != nil{
		pToBeDeleted = nil
	}
}