package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MD5加密,加密两次并转为大写
func MD5(key string) string {

	for i := 0; i < 2; i++ {
		hash := md5.New()
		hash.Write([]byte(key))
		key = strings.ToUpper(hex.EncodeToString(hash.Sum(nil)))
	}

	return key
}
