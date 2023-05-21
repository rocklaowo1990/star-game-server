package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"github.com/spf13/viper"
)

// 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func Crypto(data string, salt string) string {
	saltServer := viper.GetString("salt.server")
	return Md5Encode(saltServer + data + salt)
}
