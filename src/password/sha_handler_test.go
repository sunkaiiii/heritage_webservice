package password

import (
	"log"
	"testing"
)

func TestShaHashData(t *testing.T) {
	result := ShaHashData("123", []byte("123"))
	log.Println(result)
}
