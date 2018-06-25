package password

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
	"os"
)

// 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	file, _ := os.Open("./password/public.pem")
	buf := make([]byte, 1024)
	file.Read(buf)
	block, _ := pem.Decode(buf)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

//解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	file, err := os.Open("./password/private.pem")
	if err != nil {
		log.Println("cannot open PrivateKey.pem")
		return nil, err
	}
	buf := make([]byte, 1024)
	file.Read(buf)
	block, _ := pem.Decode(buf)
	if block == nil {
		println("private key error")
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		println("error")
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
