// Package short_url
// 参考 https://github.com/by-zhang/short-url/blob/master/ShortUrlGenerator.go
package short_url

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

const (
	VAL   = 0x3FFFFFFF
	INDEX = 0x0000003D
)

var metaByte = []byte("abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenShortUrl(fullUrl string) (shortUrl string) {
	shortUrl = fullUrl

	md5Var := md5.New()
	md5Var.Write([]byte(fullUrl))
	md5Len := md5Var.Sum(nil)
	haxStr := hex.EncodeToString(md5Len)

	var tempVal int64
	var tempUri []byte

	metaIndex := 1

	tempSubStr := haxStr[metaIndex*8 : (metaIndex+1)*8]
	hexVal, err := strconv.ParseInt(tempSubStr, 16, 64)

	if err != nil {
		return ""
	}

	tempVal = int64(VAL) & hexVal
	var index int64
	tempUri = []byte{}

	for i := 0; i < 6; i++ {
		index = INDEX & tempVal
		tempUri = append(tempUri, metaByte[index])
		tempVal = tempVal >> 5
	}

	shortUrl = string(tempUri)
	return
}
