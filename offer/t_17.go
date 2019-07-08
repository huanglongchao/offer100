/**
输入数字n,按顺序打印出从1到最大的n位十进制数
 */
package offer

import (
	"fmt"
)

func PrintToMaxOfNDigits_1(n int){
	number := 1
	for  i:=0 ; i< n; i++ {
		number = number * 10
	}
	for i:=1; i < number;i++{
		fmt.Printf("%v\t",i)
	}
}

func Print1ToMaxOfNDigits(n int){
	if n <= 0{
		return
	}
	number := make([]int,n)
	for !Increment(number) {
		PrintNumber(number)
	}
}
func Increment(number []int) bool{
	isOverflow := false
	nTakeOver := 0
	nLength := len(number)
	for i:= nLength - 1; i >= 0; i--{
		nSum := number[i] + nTakeOver
		if i == nLength - 1{
			nSum++
		}
		if nSum >= 10{
			if i == 0{
				isOverflow = true
			} else{
				nSum -= 10
				nTakeOver = 1
				number[i] = nSum
			}
		} else {
			number[i] = nSum
			break
		}
	}
	return isOverflow
}

func PrintNumber(number []int){
	isBeginning := true
	nLength := len(number)

	for i:=0; i < nLength; i++{
		if isBeginning && number[i] != 0{
			isBeginning = false
		}
		if !isBeginning{
			fmt.Printf("%v",number[i])
		}
	}
	fmt.Printf("\t")
}

func Print1ToMaxOfNDigits2(n int){
	if n <= 0{
		return
	}
	number := make([]int,n)
	for i:=0; i < 10; i++{
		number[0] = i
		Print1ToMaxOfNDigitsRecursively(number,0)
	}
}

func Print1ToMaxOfNDigitsRecursively(number []int,index int){
	if index == len(number)-1 {
		PrintNumber(number)
		return
	}
	for i:=0; i < 10; i++{
		number[index+1] = i
		Print1ToMaxOfNDigitsRecursively(number,index+1)
	}
}