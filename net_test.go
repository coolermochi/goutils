package goutils

import "testing"

func TestNet(t *testing.T) {
 	check, err := Net.Ping("google.co.jp:80")
	if err != nil {
		t.Fatalf("error %+v", err)
	}
	if !check {
		t.Fatal("error")
	}
	t.Logf("chack %v", check)

	check, err = Net.Ping("hoge:80")
	if err != nil {
		t.Fatalf("error %+v", err)
	}
	if check {
		t.Fatal("error")
	}

	t.Logf("chack %v", check)

}
