package password

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

const ID_REPEAT_COUNT = 5
const DIVIDE_NUMBER = 7
const TIME_NUMBER = 13

func ShaHashData(userName string, noEncryptPassword []byte) string {
	addSaltPassword := string(noEncryptPassword)
	passwordLen := len(noEncryptPassword)
	timeNumber := passwordLen % 7
	for i := 0; i < ID_REPEAT_COUNT; i++ {
		addSaltPassword += strconv.Itoa(timeNumber * i)
	}
	addSaltPassword += userName
	shaHash := sha256.New()
	result := shaHash.Sum([]byte(addSaltPassword))
	return hex.EncodeToString(result)
}
