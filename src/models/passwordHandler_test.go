package models

import (
	"testing"
)

func TestAesDecrpt(t *testing.T) {
	tstString := "123213213"
	testResult, err := AesEncrypt([]byte(tstString))
	println(string(testResult))
	if err != nil {
		t.Error("encrtpyError")
	}
	encryptResult, err := AesDecrypt(testResult)
	println(string(encryptResult))
	if err != nil && string(encryptResult) != tstString {
		t.Error("decrptError")
	}
}
