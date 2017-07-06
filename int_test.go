package goutils

import (
	"testing"
)

func TestNumNVL(t *testing.T) {
	i := Int.NVL(1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Int.NVL(0, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}

	i = Int.NVL(-1, 999)
	if i != -1 {
		t.Fatalf("error %d", i)
	}
}

func TestNumIIf(t *testing.T) {
	i := Int.IIf(1 == 1, 1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Int.IIf(1 == 2, 1, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}
}
