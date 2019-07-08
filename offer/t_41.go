/**
 * 数据流中的中位数
 */
package offer

import (
	heap2 "container/heap"
	"errors"
	"offer100/heap"
)

var minHeap *heap.IntMinHeap
var maxHeap *heap.IntMaxHeap

func Init(){
	minHeap = &heap.IntMinHeap{}
	maxHeap = &heap.IntMaxHeap{}
	heap2.Init(minHeap)
	heap2.Init(maxHeap)
}

func Insert(num int){
	if ( minHeap.Len() + maxHeap.Len() ) & 1 == 0{
		if maxHeap.Len() > 0 && num < (*maxHeap)[0] {
			heap2.Push(maxHeap, num)
			num = (*maxHeap)[0]
			heap2.Pop(maxHeap)
		}
		heap2.Push(minHeap,num)
	} else {
		if minHeap.Len() > 0 && (*minHeap)[0] < num{
			heap2.Push(minHeap,num)
			num = (*minHeap)[0]
			heap2.Pop(minHeap)
		}
		heap2.Push(maxHeap,num)
	}
}

func GetMedian() int{
	size := minHeap.Len() + maxHeap.Len()
	if size == 0 {
		panic(errors.New("No numbers are available"))
	}
	median := 0
	if size & 1 == 1{
		median = (*minHeap)[0]
	} else {
		median = ((*maxHeap)[0] + (*minHeap)[0]) /2
	}
	return median
}