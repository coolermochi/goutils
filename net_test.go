package goutils

import "testing"

func TestNet(t *testing.T) {
 	check, err := Net.Ping("google.co.jp:80")
	if err != nil {
		t.Fatalf("error %+v", err)
	}
	t.Logf("chack %v", check)
}
