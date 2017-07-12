// str.
package goutils

import (
	"encoding/base64"
	"strconv"
	"strings"
	"fmt"
)

// str
type str struct{}

// namespace
var String = str{}

// Def.
func (str) Def(str string, def string) string {
	return String.IIf(str != "", str, def)
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

// IIf.
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
	b64 := String.EncodeB64(str)
	fmt.Printf("%s\n", b64)
	return strings.NewReplacer("+", "-", "/", "_", "=", ".").Replace(b64)
}

// DecodeB64URLSafe.
func (str) DecodeB64URLSafe(b64 string) string {
	str := strings.NewReplacer("-", "+", "_", "/", ".", "=").Replace(b64)
	return String.DecodeB64(str)
}
