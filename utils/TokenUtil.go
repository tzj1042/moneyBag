package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//根据用户id创建一个token
//cipher,userId,millis
//cipher=Md5Encrypt(userId + key)
func CreateToken(userId, key string) string {
	var buffer bytes.Buffer
	buffer.WriteString(Md5Encrypt(userId + key))
	buffer.WriteString("," + userId + ",")
	millis := time.Now().UnixNano() / 1e6
	buffer.WriteString(strconv.FormatInt(millis, 10))
	s := buffer.String()
	encoding := base64.StdEncoding
	return encoding.EncodeToString([]byte(s))
}

//解密token
func DecryptToken(token,key string) (int, error) {
	encoding := base64.StdEncoding
	str, err := encoding.DecodeString(token)
	if err != nil {
		return 0, err
	}
	return VerifyDeCodeToken(string(str),key)
}

//验证解密后的token是否合法
func VerifyDeCodeToken(deCodeToken, key string) (int, error) {
	split := strings.Split(deCodeToken, ",")
	if len(split) != 3 {
		return 0, fmt.Errorf("token长度错误")
	}
	var cipher, userId, millis string
	cipher = split[0]
	userId = split[1]
	millis = split[2]
	return VerifyTokenSplit(cipher, userId, millis, key)
}

//验证解密后拆分后的token是否合法
func VerifyTokenSplit(cipher, userId, millis, key string) (int, error) {
	if Md5Encrypt(userId+key) != cipher {
		return 0, fmt.Errorf("token非法")
	}
	last, err := strconv.ParseInt(millis, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("token时间非法")
	}
	lastTime := time.Unix(last/1e3, 0)
	h := time.Now().Sub(lastTime).Hours()
	if h/24 > 7 {
		return 0, fmt.Errorf("token超过七天:%s,%s", millis, lastTime.Format("2006-01-02 15:04:05"))
	}
	return strconv.Atoi(userId)
}
