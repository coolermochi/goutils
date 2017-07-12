package goutils

import (
	"testing"
)

func TestIntDef(t *testing.T) {
	i := Int.Def(1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Int.Def(0, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}

	i = Int.Def(-1, 999)
	if i != -1 {
		t.Fatalf("error %d", i)
	}
}

func TestIntIIf(t *testing.T) {
	i := Int.IIf(1 == 1, 1, 999)
	if i != 1 {
		t.Fatalf("error %d", i)
	}

	i = Int.IIf(1 == 2, 1, 999)
	if i != 999 {
		t.Fatalf("error %d", i)
	}
}
