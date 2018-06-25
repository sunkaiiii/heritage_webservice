package password

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

const ID_REPEAT_COUNT = 5

func ShaHashData(id int, userName string, noEncryptPassword []byte) string {
	addSaltPassword := string(noEncryptPassword)
	for i := 0; i < ID_REPEAT_COUNT; i++ {
		addSaltPassword += strconv.Itoa(id)
	}
	addSaltPassword += userName
	shaHash := sha256.New()
	result := shaHash.Sum([]byte(addSaltPassword))
	return hex.EncodeToString(result)
}
