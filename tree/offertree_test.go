package tree

import (
	"fmt"
	"testing"
)

func TestConstructBTree(t *testing.T) {
	preorder := []int{1,2,4,7,3,5,6,8}
	inorder := []int{4,7,2,1,5,3,8,6}
	btree := ConstructBTree(preorder,inorder)
	fmt.Println("先序遍历结果：")
	PreCatTree(btree)
	fmt.Println()
	fmt.Println("中序遍历结果：")
	InCatTree(btree)
	fmt.Println()
	fmt.Println("后序遍历结果：")
	PostTree(btree)
	fmt.Println()
}