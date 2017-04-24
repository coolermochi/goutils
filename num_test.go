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

func TestNumIIf(t *testing.T) {
	i := Num.IIf(1 == 1, 1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Num.IIf(1 == 2, 1, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}
}
