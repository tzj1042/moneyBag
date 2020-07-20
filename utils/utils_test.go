package utils

import (
	"encoding/base64"
	"fmt"
	"github.com/nosixtools/solarlunar"
	"testing"
	"time"
)

func TestGetDistance(t *testing.T) {
	//天府 104.061471,30.592201
	//世豪 104.050931,30.55563
	distance := GetDistance(30.55563, 104.050931, 30.592201, 104.061471)
	ratio := Ratio(10, 3)
	fmt.Println(distance)
	fmt.Println(ratio)
	solar := solarlunar.SolarToSimpleLuanr("2020-07-01")
	fmt.Println(solar)
	content := solar[4 :]
	fmt.Println(content)
	str := TimestampToStr(710784000000 / 1000)
	fmt.Println(str)
	content2 := str[4 :]
	fmt.Println(content2)
}

func TestDeleteSlice(t *testing.T) {
	DeleteSlice()
}

//func TestDncrypt(t *testing.T) {
//	dncrypt, err := Dncrypt("L2NMpXRnyBg7oFKmr1k7e/ZrCEwFlcrA6ZM57TwuJEA1fUrQmdbmAbRoS0gTa33rvFMDAIKt5deiZoVeZ9IydOdgr9ajpRwkXGTaOAd+ed039oc5rM9aT4E/Vv7nbZFa6Phcw75ui9ivuKdQCL6dnArvOloZG2H6Ro9tnGxEbvxOnYOvQnkLGzTOtxN/TuCrR/pB9IWqB4jB7FLmttZekQ==",
//		"fSTEAn5g4gk+nUsYCr9eFw==", "IkDg6wjIUxUjjiBkj6AImg==")
//
//	fmt.Printf("== %+v", dncrypt)
//	fmt.Println()
//	fmt.Printf("cc== %+v", err)
//}

func TestEncrypt(t *testing.T) {
	bytes := []byte("huiJinSm")
	s2, err2 := Encrypt("我12w", bytes)
	fmt.Println(s2, err2)
	encrypt, err := Decrypt(s2, bytes)
	fmt.Println(encrypt, err)

	encoding := base64.StdEncoding
	str := encoding.EncodeToString([]byte("tzj1042"))
	fmt.Println("加密", str)
	decodeString, err2 := base64.StdEncoding.DecodeString(str)
	fmt.Println("解密", string(decodeString))

	fmt.Println("加密", Md5Encrypt("tzj1042"))

	token := CreateToken("122", "tzj")
	fmt.Println(token)
}

func TestDncrypt(t *testing.T) {
	fmt.Println("a")
	i := run()
	fmt.Println("b")
	fmt.Println(i)
}

func run() int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	panic("错误")
	return 4
}

func TestWeekByMonth(t *testing.T) {
	now := time.Now().AddDate(0,0,-7)
	date := WeekByMonth(now)
	fmt.Println(date)
}

func TestParseAuthToken(t *testing.T)  {
	ParseAuthToken()
}


func TestParseJwt(t *testing.T)  {
	jwt := ParseJwt("eyJhbGciOiJIUzUxMiJ9.eyJhdWQiOiJCQVNFX1NFQ1VSSVRZfGh0dHA6Ly8xMjEuNDYuNC4yMTY6ODE4MCIsImV4cCI6MTU5NDY5MjEwNiwianRpIjoiZ1dCZTdFZldzSDRxT1RnQzRHU0M2dThhUU5HcmxsR1dzT1pQeW41VTF6WVBRZzdGIiwiaXNzIjoiaHR0cDovLzEyMS40Ni40LjIxNjo4MTgwIiwiaWF0IjoxNTk0MDg3MzA2LCJzdWIiOiJQSEFkbWluIiwiY2F0ZWdvcnkiOiJVU0VSIn0.Npf8uUdaHRlddd5anfFnVS38mJgml3A5aUuM0-1uaNbTm4JU4Ap5-r0nePEqJiphKHHhiRcb-2-7z3TRlgPG1g")
	user := jwt["sub"]
	fmt.Println(user)
}
