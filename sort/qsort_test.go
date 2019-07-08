package sort

import (
	"fmt"
	"testing"
)

func TestQsort(t *testing.T) {
	data := []int{5,4,3,1,2}
	fmt.Printf("%v\n",data)
	Qsort(data,0,len(data)-1)
	fmt.Printf("%v\n",data)
}

func TestInsertSort(t *testing.T) {
	data := []int{5,4,3,1,2}
	fmt.Printf("%v\n",data)
	InsertSort(data)
	fmt.Printf("%v\n",data)
}

func TestInsertSortWithSentry(t *testing.T) {
	data := []int{0,5,4,3,1,2}
	fmt.Printf("%v\n",data)
	InsertSortWithSentry(data)
	fmt.Printf("%v\n",data)
}

func TestBubbleSort(t *testing.T) {
	data := []int{5,4,3,1,2}
	fmt.Printf("%v\n",data)
	BubbleSort(data)
	fmt.Printf("%v\n",data)
}

func TestHeapSort(t *testing.T) {
	data := []int{5,4,3,1,2}
	fmt.Printf("%v\n",data)
	HeapSort(data)
	fmt.Printf("%v\n",data)
}

func TestMegerSort(t *testing.T) {
	data := []int{5,4,3,1,2}
	fmt.Printf("%v\n",data)

	fmt.Printf("%v\n",MergeSort(data))
}

func TestMerge(t *testing.T) {
	data1 := []int{1,3,5}
	data2 := []int{2,4,6}
	fmt.Printf("%v\n",data1)
	fmt.Printf("%v\n",data2)
	result := Merge(data1,data2)
	fmt.Printf("%v\n",result)
}


func TestShellSort(t *testing.T) {
	data := []int{5,4,3,1,2}
	fmt.Printf("%v\n",data)
	ShellSort(data)
	fmt.Printf("%v\n",data)
}