package main

import "fmt"

func main(){
	data := []int{1,3,5,7}
	fmt.Println(data)
	sub := data[0:2]
	sub[1] = 2
	fmt.Println(data)


		array := []int{10, 20, 30, 40}
		slice := make([]int, 6)
		n := copy(slice, array)
		fmt.Println(n,slice)

}
