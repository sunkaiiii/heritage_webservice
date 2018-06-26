package password

import (
	"testing"
)

func TestAesDecrpt(t *testing.T) {
	testString := ShaHashData("123", []byte("123"))
	//对密码进行加密
	testResult, err := AesEncrypt([]byte(testString))
	println(string(testResult))
	if err != nil {
		t.Error("encrtpyError")
	}
	//对密文进行解密
	encryptResult, err := AesDecrypt(testResult)
	println(string(encryptResult))
	if err != nil && string(encryptResult) != testString {
		t.Error("decrptError")
	}
}
