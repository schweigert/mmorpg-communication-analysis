package password

import (
	"crypto/sha256"
	"fmt"
)

// Sum using sha256
func Sum(password string) string {
	data := []byte("hello")
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash[:])
}
