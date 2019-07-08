package offer

import (
	"fmt"
	"testing"
)

func TestGetMedian(t *testing.T) {
	//data1 := []int{}
	data2 := []int{1,2,3}
	data3 := []int{1,2,4,6}
	//data4 := []int{1,2,3,4,5,6,7,8,9,10,9,8,7,6,5,4,3,2,1}
	//data5 := []int{1,1,2,3,3,5,5}

	Init()
	for _,v := range data2{
		Insert(v)
	}
	if GetMedian() != 2{
		t.Fail()
	}
	fmt.Println(GetMedian())

	Init()
	for _,v := range data3{
		Insert(v)
	}
	if GetMedian() != 3{
		t.Fail()
	}
	fmt.Println(GetMedian())
}
