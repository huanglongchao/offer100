/**
 * 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分
 */
package offer

type F func(int) bool

func Reorder(pData []int,f F){
	if len(pData) == 0{
		return
	}
	pBegin := 0
	pEnd := len(pData) - 1
	for pBegin < pEnd{
		for pBegin < pEnd && !f(pData[pBegin]){
			pBegin++
		}
		for pBegin < pEnd && f(pData[pEnd]){
			pEnd--
		}
		if pBegin < pEnd {
			tmp  := pData[pBegin]
			pData[pBegin] = pData[pEnd]
			pData[pEnd] = tmp
		}
	}
}

func isEven(n int) bool{
	return (n&1) == 0
}