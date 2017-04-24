// str.
package goutils

import (
	"encoding/base64"
	"strconv"
	"strings"
)

// str
type str struct{}

// namespace
var Str = str{}

// NVL.
func (str) NVL(str string, def string) string {
	return Str.IIf(str != "", str, def)
}

// Atoi.
func (str) Atoi(str string) int {
	if str == "" {
		return 0
	}

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// If.
// 3項演算子.
func (str) IIf(flg bool, str1 string, str2 string) string {
	if flg {
		return str1
	} else {
		return str2
	}
}

// EncodeB64.
func (str) EncodeB64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// DecodeB64.
func (str) DecodeB64(b64 string) string {
	data, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return ""
	}
	return string(data)
}

// EncodeB64URLSafe.
func (str) EncodeB64URLSafe(str string) string {
	b64 := Str.EncodeB64(str)
	return strings.NewReplacer("=", "-", "/", "_", "+", ".").Replace(b64)
}

// DecodeB64URLSafe.
func (str) DecodeB64URLSafe(b64 string) string {
	str := strings.NewReplacer("-", "=", "_", "/", ".", "+").Replace(b64)
	return Str.DecodeB64(str)
}
