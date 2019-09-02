package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
	"fmt"
)

// 检查用户名是否合法，只接受数字和字母，且第一位非数字
func CheckUsername(username string) bool {
	if username[0] >= '0' && username[0] <= '9' {
		return false
	}
	for i := 0; i < len(username); i++ {
		if !(username[i] == '_' ||
			(username[i] >= '0' && username[i] <= '9') ||
			(username[i] >= 'a' && username[i] <= 'z') ||
			(username[i] >= 'A' && username[i] <= 'Z')) {
			return false
		}
	}
	return true
}

// 随机string
func RandString(length int) string {
	// 生成GUID(全局唯一标识符)
	buf := make([]byte, 32)
	var guid string
	if _, err := io.ReadFull(rand.Reader, buf); err != nil{
		guid = ""
	} else {
		guid = base64.URLEncoding.EncodeToString(buf)
	}

	// 截取指定的length长度的子串
	rs := []rune(guid) // 强制转成int64类型
	rsLen := len(rs)
	begin := 0
	if begin >= rsLen {
		begin = rsLen
	}
	end := begin + length
	if end > rsLen {
		end = rsLen
	}
	return string(rs[begin:end])
}

// Md5加密，返回加密后的字符串
func Md5(value string) string {
	h := md5.New()
	h.Write([]byte(value))
	return fmt.Sprintf("%s", hex.EncodeToString(h.Sum(nil)))
}

// 对字符串进行sha1计算
func Sha1(value string) string {
	s := sha1.New()
	io.WriteString(s, value)
	return fmt.Sprintf("%x", s.Sum(nil))
}