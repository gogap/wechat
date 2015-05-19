package jssdk

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

func GetNoncestrAndTimestamp() (noncestr string, timestampStr string, timestamp int64) {
	timestamp = time.Now().Unix()
	timestampStr = strconv.FormatInt(timestamp, 10)
	timestampArr := md5.Sum([]byte(timestampStr))
	hashStr := hex.EncodeToString(timestampArr[:])

	// 从md5中截取随机字符串
	noncestr = ""
	flag := true
	for _, v := range hashStr {
		if flag {
			flag = false
			continue
		}
		flag = true
		noncestr = noncestr + string(v)
	}

	return
}
