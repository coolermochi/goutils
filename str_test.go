package goutils

import (
	"testing"
)

func TestStrNVL(t *testing.T) {
	s := Str.NVL("a", "b")
	if s != "a" {
		t.Fatalf("error %s", s)
	}

	s = Str.NVL("", "b")
	if s != "b" {
		t.Fatalf("error %s", s)
	}

	s = Str.NVL("", "")
	if s != "" {
		t.Fatalf("error %s", s)
	}
}

func TestStrAtoi(t *testing.T) {
	i := Str.Atoi("123")
	if i != 123 {
		t.Fatalf("error %d", i)
	}

	i = Str.Atoi("abc")
	if i != 0 {
		t.Fatalf("error %d", i)
	}

	i = Str.Atoi("")
	if i != 0 {
		t.Fatalf("error %d", i)
	}

}

func TestStrIIf(t *testing.T) {
	s := Str.IIf(1 == 1, "a", "b")
	if s != "a" {
		t.Fatalf("error %s", s)
	}

	s = Str.IIf(1 == 2, "a", "b")
	if s != "b" {
		t.Fatalf("error %s", s)
	}
}

func TestStrB64(t *testing.T) {
	text := "hoge"
	enc := Str.EncodeB64(text)
	dec := Str.DecodeB64(enc)
	if text != dec {
		t.Fatalf("error text:%s enc:%s dec:%s", text, enc, dec)
	}

	t.Logf("text:%s enc:%s dec:%s", text, enc, dec)
}

func TestStrB64URLSafe(t *testing.T) {
	text := "hoge"
	enc := Str.EncodeB64URLSafe(text)
	dec := Str.DecodeB64URLSafe(enc)
	if text != dec {
		t.Fatalf("error text:%s enc:%s dec:%s", text, enc, dec)
	}

	t.Logf("text:%s enc:%s dec:%s", text, enc, dec)
}
