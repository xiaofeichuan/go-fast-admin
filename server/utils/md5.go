package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
