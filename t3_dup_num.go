//  题目3:找出数组中重复的数字
//  在一个长度为n的数组里的所有数字都在0-n-1范围内，数组中某些元素是重复的，
//  请找出数组中任意一个重复的数字，例如，如果输入长度为7的数组{2,3,1,0,2,5,3}，那么对应的输出是2或者3
package main

import (
   "fmt"
)

func dup(nums []int , len int, duplication *int) bool{
    if nums == nil || len <=0 {
    	return false
    }
    for i:=0; i<len; i++ {
    	if nums[i] < 0 || nums[i] > len - 1{
    		return false
    	}
    }
    for i:=0; i<len; i++ {
    	for nums[i] != i {
    		if nums[i] == nums[nums[i]]{
    			*duplication = nums[i]
    			return true
    		}
    		tmp = nums[i]
    		nums[i] = nums[tmp]
    		nums[tmp] = tmp
    	}
    }
    return false
}

var nums []int = []int{2,3,1,0,2,5,3}
var tmp int

func main(){
	    dup(nums,7,&tmp)



        fmt.Printf("dup num is %d\n",tmp)
}
