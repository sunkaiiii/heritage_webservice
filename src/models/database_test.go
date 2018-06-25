package models

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"log"
	"os"
	"testing"
)

func TestRsaDecrypt(t *testing.T) {
	sql := "select ID,USER_PASSWORD from user_info"
	rows, err := DB.Query(sql)
	defer rows.Close()
	if err != nil {
		t.Error(err.Error())
	}
	var id int
	var password string
	for rows.Next() {
		rows.Scan(&id, &password)
		passwordByte, _ := base64.StdEncoding.DecodeString(password)
		encryptPassword, _ := RsaDecrypt(passwordByte)
		print(string(encryptPassword))
	}
}

//解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	file, err := os.Open("../password/private.pem")
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
