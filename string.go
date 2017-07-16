// str.
package goutils

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"bytes"
	//"strings"
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
		fmt.Printf("error DecodeB64 %s\n", err.Error())
		return ""
	}
	return string(data)
}

// EncodeB64URLSafe.
func (str) EncodeB64URLSafe(str string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(str))
}

// DecodeB64URLSafe.
func (str) DecodeB64URLSafe(b64 string) string {
	dec := base64.NewDecoder(base64.RawURLEncoding, bytes.NewBufferString(b64))
	buf := new(bytes.Buffer)
	buf.ReadFrom(dec)
	return buf.String()
}
