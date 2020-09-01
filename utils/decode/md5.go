package decode

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为 123456
	res := hex.EncodeToString(h.Sum(nil))
	return res
}
