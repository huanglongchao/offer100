package offer

import (
	"fmt"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	data := []int{5,3,4,2,1}
	fmt.Println(GetLeastNumbers(data,3))
}

func TestGetLeastNumbers2(t *testing.T) {
	data := []int{5,3,4,2,1}
	fmt.Println(GetLeastNumbers2(data,3))
}