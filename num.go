// num.
package goutils

// num
type num struct{}

// namespace
var Num = num{}

// NVL.
func (num) NVL(i int, def int) int {
	return Num.IIf(i != 0, i, def)
}

// If.
// 3項演算子
func (num) IIf(flg bool, i1 int, i2 int) int {
	if flg {
		return i1
	} else {
		return i2
	}
}
