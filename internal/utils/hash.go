package utils

import (
	"crypto/sha256"
	"fmt"
)

func GetHashedString(str string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(str)))
}
