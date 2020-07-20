package utils

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 获取IP地址
func GetIp(r *http.Request) string {
	ipadd := r.RemoteAddr
	sArr := strings.Split(ipadd, ":")
	return sArr[0]
}

//md5加密
func Md5Encrypt(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

//DES加密
func Encrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

//DES解密
func Decrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

//类型断言
func Typeof(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return "unknown"
	}
}

//根据json对象字符串生成结构体
func createStruct(jsonStr string, structName string) {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println("转化错误:", err)
	}
	var buffer bytes.Buffer
	buffer.WriteString("type ")
	buffer.WriteString(structName)
	buffer.WriteString(" struct {\n")
	for k, v := range m {
		runes := []rune(k)
		buffer.WriteString(strings.ToUpper(string(runes[0])))
		buffer.WriteString(string(runes[1:]))
		buffer.WriteString("   ")
		buffer.WriteString(Typeof(v))
		buffer.WriteString("     `json:\"")
		buffer.WriteString(k)
		buffer.WriteString("\"`")
		buffer.WriteString("\n")
	}
	buffer.WriteString("}")
	fmt.Println(buffer.String())
}

// 获取用户授权token
func GetUserToken(userId int64, myKey string) (token string, err error) {
	var str = myKey + "." + strconv.FormatInt(time.Now().Unix(), 10) + "." + strconv.FormatInt(userId, 10)
	return Encrypt(str, []byte(myKey))
}

// 验证token
// encryptToken 加密的token
// myKey 密钥
// 返回正确的用户id
func VerifyToken(encryptToken string, myKey string) (userId int, err error) {
	var b = []byte(myKey)
	var token string
	token, err = Decrypt(encryptToken, b)
	if err != nil {
		fmt.Println("解密错误：", err)
		return 0, fmt.Errorf("解密失败，token非法:%s", err)
	}
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return 0, fmt.Errorf("解析失败，token非法:%s", "len is not 3")
	}
	var last int64
	last, err = strconv.ParseInt(split[1], 10, 64)
	lastTime := time.Unix(last, 0)
	m := time.Now().Sub(lastTime).Hours()
	if m/24 > 7 {
		return 0, fmt.Errorf("token超过七天:%s,%s", token, lastTime.Format("2006-01-02 15:04:05"))
	}
	if err != nil {
		fmt.Println("加密密错误：", err)
		return 0, fmt.Errorf("解密失败，token非法:%s", err)
	}
	return strconv.Atoi(split[2])
}

// 生成指定位数的随机数
func RandomNo(n int) string {
	var buffer bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < n; i++ {
		buffer.WriteString(strconv.Itoa(r.Intn(10)))
	}
	return buffer.String()
}

// 生成指定位数的随机数的订单号
func GetOrderNo(n int) string {
	no := RandomNo(n)
	return GetNowString("20060102150405") + no
}

//生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	b := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

type src struct {
	Id   string
	Name string
	Code int
}

type Target struct {
	Name string
	Code int
	Id   int
}

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	bytes, _ := json.Marshal(DstStructPtr)
	json.Unmarshal(bytes, SrcStructPtr)
}

// src 值类型
// tar指针类型
func StructCopy2(src interface{}, tar interface{}) {
	getType := reflect.TypeOf(src)
	getValue := reflect.ValueOf(src)
	getValue2 := reflect.ValueOf(tar)
	elem := getValue2.Elem()
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i)
		//fmt.Printf("1名:%s 类型:%v = 值:%v\n", field.Name, field.Type, value.Interface())
		if getValue2.Kind() == reflect.Ptr {
			for i := 0; i < elem.NumField(); i++ {
				fieldInfo := elem.Type().Field(i)
				//fmt.Printf("2名:%s 类型:%v = 值:%v\n", fieldInfo.Name, fieldInfo.Type, getValue2.Elem().FieldByName(fieldInfo.Name))
				name := elem.FieldByName(fieldInfo.Name)
				if field.Name == fieldInfo.Name && name.Kind() == field.Type.Kind() { /*
					switch name.Kind() {
					case reflect.String:
						*(*string)(unsafe.Pointer(name.Addr().Pointer())) = value.(string)
					case reflect.Int64:
						*(*int64)(unsafe.Pointer(name.Addr().Pointer())) = value.(int64)
					case reflect.Int:
						*(*int)(unsafe.Pointer(name.Addr().Pointer())) = value.(int)
					case reflect.Bool:
						*(*bool)(unsafe.Pointer(name.Addr().Pointer())) = value.(bool)
					case reflect.Float64:
						*(*float64)(unsafe.Pointer(name.Addr().Pointer())) = value.(float64)
					}*/
					elem.FieldByName(field.Name).Set(value)
					break
				}
			}
		}
	}
}

// 两个坐标的距离
// 返回值的单位为千米
func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := float64(6371000) // 6378137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / 1000
}

//结构体的反射遍历(值类型)
func StructFor(src interface{}) {
	getType := reflect.TypeOf(src)
	getValue := reflect.ValueOf(src)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
}

//结构体的反射遍历(指针类型)
func StructPtrFor(src interface{}) {
	getValue2 := reflect.ValueOf(src)
	if getValue2.Kind() == reflect.Ptr {
		elem := getValue2.Elem()
		for i := 0; i < elem.NumField(); i++ {
			fieldInfo := elem.Type().Field(i)
			fmt.Println("名：", fieldInfo.Name)
			fmt.Println("类型：", fieldInfo.Type)
			fmt.Println("值：", getValue2.Elem().FieldByName(fieldInfo.Name))
		}
	}
}

var (
	w sync.RWMutex
)

/*func main() {
json:=`{
      "partner_order_id": "2019060411054412811",
      "full_name": "Crowd Technology Corp.",
      "partner_name": "Hui玩app",
      "channel": "AlipayOnline",
      "sdk_params": "service=\"mobile.securitypay.pay\"&partner=\"2088031415022202\"&_input_charset=\"utf-8\"&product_code=\"NEW_WAP_OVERSEAS_SELLER\"&notify_url=\"https://pay.alphapay.ca/api/v1.0/alipay/online_orders/R9AMAU-20190603190543989-YWLFDOU/notify\"&out_trade_no=\"R9AMAU-20190603190543989-YWLFDOU\"&subject=\"HUI玩支付2019060411054412811\"&payment_type=\"1\"&seller_id=\"2088031415022202\"&currency=\"CAD\"&rmb_fee=\"10.00\"&body=\"HUI玩支付\"&forex_biz=\"FP\"&it_b_pay=\"2019-06-06 02:05:43\"&secondary_merchant_id=\"R9AMAU\"&secondary_merchant_name=\"Hui玩app\"&secondary_merchant_industry=\"5311\"&sign=\"qO8SaClXCh8gT1tExn9N3a%2FiV5MM8N0QgroOxZgplxYkPPOUgRHbxchX9IK8nNocVZ%2B9zz4FvjwAfD191kuriIrUdfa1Ql0B45hfUT2KyKfyhYcdftABns41FWeIZDTbaTwZZ%2FBfax7kQDloEf7n15to2W15ouA0QhHPNlOkmzw%3D\"&sign_type=\"RSA\"",
      "result_code": "SUCCESS",
      "partner_code": "R9AMAU",
      "order_id": "R9AMAU-20190603190543989-YWLFDOU",
      "return_code": "SUCCESS"
    }`
    createStruct(json,"PayInfo")
}*/

func DeleteSlice() {
	s := []int{1, 2, 3, 4, 5}
	for i, v := range s {
		if v == 3{
			s = append(s[:i], s[i+1:]...)
		}
	}
	for _, v := range s {
		fmt.Println(v)
	}
}

func ParseAuthToken()  {
	mySignKey := "mySecret"     //密钥，同java代码
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(mySignKey)   //需要用和加密时同样的方式转化成对应的字节数组
	if err != nil {
		fmt.Println("base64 decodeString failed.", err)
		return
	}
	token:="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDQ3NDk0NDIsImZvbyI6ImJhciIsImlhdCI6MTU0NDc1MzA0Mn0.t1NwZOUJP4Vj3L4YAiHletRNlc8vEtOMrjRAiyKl8aA"
	parseAuth, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return mySignKeyBytes, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println(parseAuth.Claims)
}

//解析tokenString
func ParseJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte("mySecret"), nil
	})

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
	return claims
}