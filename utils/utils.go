package utils

import (
	"os"
	"crypto/md5"
	"io"
	"encoding/hex"
)

// MD5 File
// 用于检查文件是否变动
func Md5File(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "",err
	}
	h := md5.New()
	_, err = io.Copy(h,file)
	if err != nil {
		return "",err
	}
	return hex.EncodeToString(h.Sum([]byte{})), nil
}





