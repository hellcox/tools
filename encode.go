package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"hash/crc32"
	"io"
	"strconv"
)

const (
	MD5    = "md5"
	BASE64 = "base64"
	CRC32  = "crc32"
)

func encode(action, val, password string) string {
	switch action {
	case MD5:
		return md5Encode(val)
	case BASE64:
		return base64Encode(val)
	case CRC32:
		res := crc32Encode(val)
		return strconv.FormatUint(uint64(res), 10)
	default:
		return "类型错误"
	}
	return ""
}

func decode(action, val, password string) string {
	if action == MD5 {
		return "不支持解密：md5"
	}

	switch action {
	case MD5:
		return "不支持 MD5 Decode"
	case BASE64:
		res, err := base64Decode(val)
		if err != nil {
			return "解析失败：" + err.Error()
		}
		return res
	case CRC32:
		return "不支持 CRC32 Decode"
	default:
		return "类型错误"
	}
	return ""
}

func md5Encode(val string) string {
	w := md5.New()
	_, _ = io.WriteString(w, val)            //将str写入到w中
	md5str2 := fmt.Sprintf("%x", w.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式
	return md5str2
}

func base64Encode(val string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(val))
	return sEnc
}

func base64Decode(val string) (string, error) {
	sDec, err := base64.StdEncoding.DecodeString(val)
	return string(sDec), err
}

func crc32Encode(val string) uint32 {
	return crc32.ChecksumIEEE([]byte(val))
}
