package password

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"log"
	"os"
)

var ivDefValue string
var key []byte

func init() {
	file, err := os.Open("./password/aes_block_key.key")
	if err != nil {
		log.Println(err.Error())
		return
	}
	buf := make([]byte, 128)
	n, err := file.Read(buf)
	if err != nil {
		log.Println(err.Error())
	}
	ivDefValue = string(buf[0:n])
	file.Close()
	file, err = os.Open("./password/aes_key.key")
	if err != nil {
		log.Println(err.Error())
		return
	}
	buf = make([]byte, 256)
	n, err = file.Read(buf)
	key = buf[0:n]
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func AesEncrypt(plaintext []byte) ([]byte, error) {
	if len(ivDefValue) == 0 || key == nil {
		return nil, errors.New("no key")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}
	blockSize := block.BlockSize()
	plaintext = pkcS5Padding(plaintext, blockSize)
	iv := []byte(ivDefValue)
	blockMode := cipher.NewCBCEncrypter(block, iv)

	ciphertext := make([]byte, len(plaintext))
	blockMode.CryptBlocks(ciphertext, plaintext)

	return ciphertext, nil
}

func AesDecrypt(ciphertext []byte) ([]byte, error) {
	if len(ivDefValue) == 0 && key == nil {
		return nil, errors.New("no key")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.New("invalid decrypt key")
	}

	blockSize := block.BlockSize()

	if len(ciphertext) < blockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := []byte(ivDefValue)
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	blockModel := cipher.NewCBCDecrypter(block, iv)

	plaintext := make([]byte, len(ciphertext))
	blockModel.CryptBlocks(plaintext, ciphertext)
	plaintext = pkcS5UnPadding(plaintext)

	return plaintext, nil
}

func pkcS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
