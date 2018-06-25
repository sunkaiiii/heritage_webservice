package password

import (
	"log"
	"testing"
)

func TestShaHashData(t *testing.T) {
	result := ShaHashData(1, "123", []byte("123"))
	log.Println(result)
}
