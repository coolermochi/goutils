package goutils

import (
	"testing"
)

func TestStringDef(t *testing.T) {
	s := String.Def("a", "b")
	if s != "a" {
		t.Fatalf("error %s", s)
	}

	s = String.Def("", "b")
	if s != "b" {
		t.Fatalf("error %s", s)
	}

	s = String.Def("", "")
	if s != "" {
		t.Fatalf("error %s", s)
	}
}

func TestStringAtoi(t *testing.T) {
	i := String.Atoi("123")
	if i != 123 {
		t.Fatalf("error %d", i)
	}

	i = String.Atoi("abc")
	if i != 0 {
		t.Fatalf("error %d", i)
	}

	i = String.Atoi("")
	if i != 0 {
		t.Fatalf("error %d", i)
	}

}

func TestStringIIf(t *testing.T) {
	s := String.IIf(1 == 1, "a", "b")
	if s != "a" {
		t.Fatalf("error %s", s)
	}

	s = String.IIf(1 == 2, "a", "b")
	if s != "b" {
		t.Fatalf("error %s", s)
	}
}

func TestStringB64(t *testing.T) {
	text := "hoge"
	enc := String.EncodeB64(text)
	dec := String.DecodeB64(enc)
	if text != dec {
		t.Fatalf("error text:%s enc:%s dec:%s", text, enc, dec)
	}

	t.Logf("text:%s enc:%s dec:%s", text, enc, dec)
}

func TestStringB64URLSafe(t *testing.T) {
	text := "hoge"
	enc := String.EncodeB64URLSafe(text)
	dec := String.DecodeB64URLSafe(enc)
	if text != dec {
		t.Fatalf("error text:%s enc:%s dec:%s", text, enc, dec)
	}

	t.Logf("text:%s enc:%s dec:%s", text, enc, dec)
}
