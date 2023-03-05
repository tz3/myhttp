package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash returns the MD5 hash of the specified byte slice
func MD5Hash(data []byte) string {
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
