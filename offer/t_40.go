/**
 * 最小的K个数
 * 输入n个整数，找出其中自小的k个数
 */
package offer

import "offer100/sort"

func GetLeastNumbers(data []int, k int) (result []int){

	if data == nil || k <=0 || k > len(data){
		return nil
	}
	start := 0
	end := len(data) -1
	index := sort.Partition(data,start,end)
	for index != k - 1{
		if index > k-1{
			end = index - 1
			index = sort.Partition(data,start,end)
		} else {
			start = index + 1
			index = sort.Partition(data,start,end)
		}
	}
	return data[0:k]
}
func GetLeastNumbers2(data []int, k int) (result []int){
	if data == nil || k <= 0 || k > len(data){
		return nil
	}
	length := len(data)
	//建大小为k的堆
	for i := k/2 - 1; i >=0; i--{
		sort.HeapAdjust(data,i,k-1)
	}
	for i:=k; i < length; i++{
		if data[i] < data[0]{
			data[0] = data[i]
			sort.HeapAdjust(data,0,k-1)
		}
	}
	return data[:k]
}