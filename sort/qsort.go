package sort
/**
 * 快速排序
 */
func Qsort(data []int, low int, high int){
	if low < high{
		pivot := Partition(data,low,high)
		Qsort(data,low,pivot-1)
		Qsort(data,pivot+1,high)
	}
}
func Partition(data []int, low int, high int) int{
	pivot := data[low]
	for low < high{
		for low < high && data[high] >= pivot{
			high--
		}
		data[low] = data[high]
		for low < high && data[low] <= pivot{
			low++
		}
		data[high] = data[low]
	}
	data[low] = pivot
	return low
}
/**
 * 插入排序（不带哨兵）
 */
func InsertSort(data []int){
	for i:=1; i < len(data); i++{
		if data[i] < data[i-1]{
			tmp := data[i]
			var j int
			for j=i-1; j >=0 && data[j] > tmp; j--{
				data[j+1] = data[j]
			}
			data[j+1] = tmp
		}
	}
}
/**
 * 插入排序（带哨兵）
 */
func InsertSortWithSentry(data []int){
	for i:=2; i < len(data); i++{
		if data[i] < data[i-1]{
			data[0] = data[i]
			data[i] = data[i-1]
			var j int
			for j=i-2; data[j] > data[0]; j--{
				data[j+1] = data[j]
			}
			data[j+1] = data[0]
		}
	}
}

/**
 * 希尔排序
 */

/**
 * 冒泡排序
 */
func BubbleSort(data []int){
	len := len(data)
	for i:=0; i<len; i++{
		for j:=0; j < len-i-1; j++{
			if data[j+1] < data[j]{
				data[j],data[j+1] = data[j+1],data[j]
			}
		}
	}
}
/**
 * 归并排序
 */
func Merge(left,right []int)(result []int){
	l,r := 0,0
	for l < len(left) && r < len(right){
		if left[l] < right[r]{
			result = append(result,left[l])
			l++
		} else{
			result = append(result,right[r])
			r++
		}
	}
	result = append(result,left[l:]...)
	result = append(result,right[r:]...)

	return result
}
func MergeSort(data []int) []int{
	length := len(data)
	if length <= 1{
		return data
	}
	num := length/2
	left := MergeSort(data[:num])
	right := MergeSort(data[num:])
	return Merge(left,right)
}
/**
 * 堆排序
		0
	1		2
  3	  4	  5   6
上图len = 7
注意几个关键点
非叶子节点为    len/2 - 1至0
左孩子节点		j=2*i+1
有孩子节点		j+1
大顶堆
 */
func HeapAdjust(data []int, start int, end int){
	tmp := data[start]
	for j:=2*start+1; j<=end ; j=2*j+1  {
		if j<end && data[j] < data[j+1]{
			j++
		}
		if tmp > data[j]{
			break
		}
		data[start] = data[j]
		start = j
	}
	data[start] = tmp
}
//小顶堆调整
func HeapAdjust2(data[] int, start int, end int){
	tmp := data[start]
	for j:=2*start+1; j <=end; j = 2*j+1{
		if j < end && data[j] > data[j+1]{
			j++
		}
		if tmp < data[j]{
			break
		}
		data[start] = data[j]
		start = j
	}
	data[start] = tmp
}
func HeapSort(data []int){
	//建堆
	for i:= len(data)/2-1; i>=0; i--{
		HeapAdjust(data,i,len(data)-1)
	}
	for i:= len(data)-1; i>0; i--{
		data[0],data[i] = data[i],data[0]
		HeapAdjust(data,0,i-1)
	}
}

/**
 * 希尔排序
 */
func ShellSort(data []int){
	for step := len(data)/2; step > 0; step /= 2{
		for i:=step; i< len(data); i++{
			for j:=i-step; j>=0 && data[j+step] < data[j]; j -= step{
				data[j],data[j+step] = data[j+step],data[j]
			}
		}
	}
}