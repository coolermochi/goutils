package goutils

import (
	"testing"
)

func TestCryptPlain(t *testing.T) {

	typo := CryptPlain
	secret := "hoge"
	text := "abcdefg"

	// 暗号化
	enc := Crypt.Encrypt(typo, secret, text)
	if enc == "" {
		t.Errorf("error encypt secret:%s text:%s", secret, text)
	}

	// 復号化
	dec := Crypt.Decrypt(typo, secret, enc)
	if dec == "" {
		t.Errorf("error decrypt secret:%s text:%s enc:%s", secret, text, enc)
	}

	// check
	if text != dec {
		t.Errorf("error check secret:%s text:%s enc:%s dec:%s", secret, text, enc, dec)
	}

	t.Logf("[%s] -> [%s] -> [%s]", text, enc, dec)
}

func TestCryptB64(t *testing.T) {

	typo := CryptB64
	secret := "hoge"
	text := "abcdefg"

	// 暗号化
	enc := Crypt.Encrypt(typo, secret, text)
	if enc == "" {
		t.Errorf("error encypt secret:%s text:%s", secret, text)
	}

	// 復号化
	dec := Crypt.Decrypt(typo, secret, enc)
	if dec == "" {
		t.Errorf("error decrypt secret:%s text:%s enc:%s", secret, text, enc)
	}

	// check
	if text != dec {
		t.Errorf("error check secret:%s text:%s enc:%s dec:%s", secret, text, enc, dec)
	}

	t.Logf("[%s] -> [%s] -> [%s]", text, enc, dec)
}

func TestCryptB64URLSafe(t *testing.T) {

	typo := CryptB64URLSafe
	secret := "hoge"
	text := "abcdefg"

	// 暗号化
	enc := Crypt.Encrypt(typo, secret, text)
	if enc == "" {
		t.Errorf("error encypt secret:%s text:%s", secret, text)
	}

	// 復号化
	dec := Crypt.Decrypt(typo, secret, enc)
	if dec == "" {
		t.Errorf("error decrypt secret:%s text:%s enc:%s", secret, text, enc)
	}

	// check
	if text != dec {
		t.Errorf("error check secret:%s text:%s enc:%s dec:%s", secret, text, enc, dec)
	}

	t.Logf("[%s] -> [%s] -> [%s]", text, enc, dec)
}
