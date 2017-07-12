// integer.
package goutils

// integer
type integer struct{}

// namespace
var Int = integer{}

// Def.
// 0だった場合にデフォルト値返却
func (integer) Def(i int, def int) int {
	return Int.IIf(i != 0, i, def)
}

// IIf.
// 3項演算子
func (integer) IIf(flg bool, i1 int, i2 int) int {
	if flg {
		return i1
	} else {
		return i2
	}
}
