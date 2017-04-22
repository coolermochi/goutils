package goutils

import (
	"testing"
)

func TestNumNVL(t *testing.T) {
	i := Num.NVL(1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Num.NVL(0, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}

	i = Num.NVL(-1, 999)
	if i != -1 {
		t.Fatalf("error %d", i)
	}
}

func TestNumIf(t *testing.T) {
	i := Num.If(1 == 1, 1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Num.If(1 == 2, 1, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}
}
