package goutils

import (
	"testing"
)

func TestSliceSum(t *testing.T) {

	nums := []int{1,2,3,4,5,6,7,8,9,10}

	sum := Slice.Sum(nums)
	if sum != 55 {
		t.Fatalf("error sum %d", sum)
	}
}
