package offer

import (
	"fmt"
	"testing"
)

func TestReorder(t *testing.T) {
	pdata := []int{2,4,6,1,3,5}
	fmt.Printf("%v\n",pdata)
	Reorder(pdata,isEven)
	fmt.Printf("%v\n",pdata)
}