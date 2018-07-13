package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// GetCSRFToken 生成CSRF Token
func GetCSRFToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	io.WriteString(h, "abcdefghijklmnopgrstuvwxyzxxxxxx")
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
