// slicer.
package goutils

// slicer
type slicer struct{}

// namespace
var Slice = slicer{}

// Sum.
func (slicer) Sum(nums []int) int {
	sum := 0
	if nums == nil || len(nums) == 0 {
		return sum
	}

	for _, n := range nums {
		sum += n
	}

	return sum
}
