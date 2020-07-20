package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

const (
	YYYYMMddHHmmss = "2006-01-02 15:04:05"
	YYYYMMddHHmm   = "2006-01-02 15:04"
	YYYYMMdd       = "2006-01-02"
)

type JsonDate time.Time
type JsonTime time.Time
type Time time.Time

func (p *JsonDate) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(YYYYMMdd, string(data), time.Local)
	*p = JsonDate(local)
	return err
}

func (p *JsonTime) UnmarshalJSON(data []byte) error {
	local, err := time.ParseInLocation(YYYYMMddHHmmss, string(data), time.Local)
	*p = JsonTime(local)
	return err
}

func (c JsonDate) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, YYYYMMdd)
	data = append(data, '"')
	return data, nil
	/*stamp := fmt.Sprintf("%d", time.Time(c).Unix())
	return []byte(stamp), nil*/
}

func (c JsonTime) MarshalJSON() ([]byte, error) {
	data := make([]byte, 0)
	data = append(data, '"')
	data = time.Time(c).AppendFormat(data, YYYYMMddHHmmss)
	data = append(data, '"')
	return data, nil
}

func (c JsonDate) String() string {
	return time.Time(c).Format(YYYYMMdd)
}

func (c JsonTime) String() string {
	return time.Time(c).Format("2006-01-02 15:04:05")
}

func ToDate(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02", in)
	return out, err
}

func ToDatetime(in string) (out time.Time, err error) {
	out, err = time.Parse("2006-01-02 15:04:05", in)
	return out, err
}

func GetNowString1() string {
	return time.Now().Format(YYYYMMdd)
}

func GetNowString2() string {
	return time.Now().Format(YYYYMMddHHmm)
}

func GetNowString3() string {
	return time.Now().Format(YYYYMMddHHmmss)
}

func GetNowString(format string) string {
	return time.Now().Format(format)
}

func TimeTOString(t time.Time) string {
	return t.Format(YYYYMMdd)
}

//相差的分钟
func DiffMinute(t1, t2 string) (float64, error) {
	parse, err := time.Parse(YYYYMMddHHmmss, t1)
	if err != nil {
		return 0, fmt.Errorf("t1错误%s", err.Error())
	}
	parse2, err2 := time.Parse(YYYYMMddHHmmss, t2)
	if err2 != nil {
		return 0, fmt.Errorf("t2错误%s", err2.Error())
	}
	return parse2.Sub(parse).Minutes(), nil
}

func GetNowInt64() int64 {
	return time.Now().Unix()
}

func DiffHours(t1, t2 time.Time) float64 {
	return t2.Sub(t1).Hours()
}

func DiffHours2(t1, t2 string) (float64, error) {
	parse, err := time.Parse(YYYYMMdd, t1)
	if err != nil {
		return 0, fmt.Errorf("t1错误%s", err.Error())
	}
	parse2, err2 := time.Parse(YYYYMMdd, t2)
	if err2 != nil {
		return 0, fmt.Errorf("t2错误%s", err2.Error())
	}
	return parse2.Sub(parse).Hours(), nil
}

//之后多少分钟
//minute -60m之前多少分钟60m之后多少分钟
func AfterTime(timeStr, minute string) (string, error) {
	localTime, err := time.ParseInLocation(YYYYMMddHHmmss, timeStr, time.Local)
	if err != nil {
		return "", err
	}
	sm, err2 := time.ParseDuration(minute)
	if err2 != nil {
		return "", err2
	}
	return localTime.Add(sm).Format(YYYYMMddHHmmss), nil
}

func AfterTime2(timeStr time.Time, minute string) (string, error) {
	sm, err2 := time.ParseDuration(minute)
	if err2 != nil {
		return "", err2
	}
	return timeStr.Add(sm).Format(YYYYMMddHHmmss), nil
}

// MarshalJSON 序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("%d", time.Time(t).UnixNano()/1e6)
	return []byte(stamp), nil
}

// String 重写String方法
func (t *Time) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// CBC 模式
//解密
/**
* rawData 原始加密数据
* key  密钥
* iv  向量
 */
func Dncrypt(rawData, key, iv string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(rawData)
	key_b, err_1 := base64.StdEncoding.DecodeString(key)
	iv_b, _ := base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return "", err
	}
	if err_1 != nil {
		return "", err_1
	}
	dnData, err := AesCBCDncrypt(data, key_b, iv_b)
	if err != nil {
		return "", err
	}
	return string(dnData), nil
}

// 解密
func AesCBCDncrypt(encryptData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockSize := block.BlockSize()
	if len(encryptData) < blockSize {
		panic("ciphertext too short")
	}
	if len(encryptData)%blockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(encryptData, encryptData)
	// 解填充
	encryptData = PKCS7UnPadding(encryptData)
	return encryptData, nil
}

//去除填充
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// 本周一
func WeekOne() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStart
}

//本月第一天
func NowMonthOne() time.Time {
	d := time.Now()
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 当月第一天
func MonthOne(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

//获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}

// 时间戳转时间字符串
func TimestampToStr(timestamp int64) string {
	//时间戳 to 时间
	tm := time.Unix(timestamp, 0)
	return tm.Format(YYYYMMdd)
}

// 时间字符串转时间戳
func StrToTimestamp(str string) int64 {
	tm2, _ := time.Parse(YYYYMMdd, str)
	return tm2.Unix()
}

//判断时间是当年的第几周
func WeekByYear(t time.Time) int {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return week
}

//判断本月后的时间是当月的第几周
func WeekByMonth(t time.Time) int {
	//是当年的第几周
	yearWeek := WeekByYear(t)
	one := MonthOne(t)
	//当月第一天是当年第几周
	yearWeek2 := WeekByYear(one)
	// 相差多少周
	diffWeek := yearWeek - yearWeek2
	var week int
	if diffWeek == 0 {
		week = 1
	} else {
		week = diffWeek + 1
	}
	return week
}

//
//// FieldType 数据类型
//func (t *Time) FieldType() int64 {
//	return t.Time.Unix()
//
//}
//
//// SetRaw 读取数据库值
//func (t *Time) SetRaw(value interface{}) error {
//	switch value.(type) {
//	case time.Time:
//		t.Time = value.(time.Time)
//	}
//	return nil
//}
//
//// RawValue 写入数据库
//func (t *Time) RawValue() interface{} {
//	str := t.Format("2006-01-02 15:04:05")
//	if str == "0001-01-01 00:00:00" {
//		return nil
//	}
//	return str
//}
