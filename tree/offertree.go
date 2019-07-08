package tree

import (
	"fmt"
)

type BinaryTreeNode struct {
	m_nValue 	int
	m_pLeft 	*BinaryTreeNode
	m_pRight	*BinaryTreeNode
}
/**
输入某二叉树的前序遍历和中序遍历的结果，请重建改二叉树
假设输入的前序和中序遍历的结果不包含重复的数字
例如 输入前序{1,2,4,7,3,5,6,8} 中序{4,7,2,1,5,3,8,6}
 */
func ConstructBTree(preorder, inorder []int) *BinaryTreeNode{

	if preorder == nil || inorder == nil || len(preorder) == 0 {
		return nil
	}
	root := &BinaryTreeNode{
		m_nValue:preorder[0],
	}
	leftLen := 0
	rightLen := 0
	for _,v := range inorder{
		if v == root.m_nValue{
			break
		}
		leftLen++
	}
	rightLen = len(preorder) - leftLen - 1

	if leftLen > 0{
		root.m_pLeft = ConstructBTree(preorder[1:leftLen + 1],inorder[0:leftLen])
	}
	if rightLen > 0{
		root.m_pRight = ConstructBTree(preorder[leftLen+1:],inorder[leftLen+1:])
	}
	return root
}

func PreCatTree(t *BinaryTreeNode){
	fmt.Printf("%v ",t.m_nValue)
	if t.m_pLeft != nil{
		PreCatTree(t.m_pLeft)
	}
	if t.m_pRight != nil {
		PreCatTree(t.m_pRight)
	}
}

func InCatTree(t *BinaryTreeNode){
	if t.m_pLeft != nil{
		InCatTree(t.m_pLeft)
	}
	fmt.Printf("%v ",t.m_nValue)
	if t.m_pRight != nil {
		InCatTree(t.m_pRight)
	}
}

func PostTree(t *BinaryTreeNode){
	if t.m_pLeft != nil{
		PostTree(t.m_pLeft)
	}
	if t.m_pRight != nil {
		PostTree(t.m_pRight)
	}
	fmt.Printf("%v ",t.m_nValue)
}
