package utils

import (
	"math"
)

//4舍5入保留指定位数
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

//保留两位小数，去尾法
func Decimal(value float64) float64 {
	return math.Trunc(value*1e2) * 1e-2
}

//计算占比
//son占sum的多少
func Ratio(son, sum float64) float64 {
	n10 := math.Pow10(2)
	return math.Trunc(((son/sum)+0.5/n10)*n10) / n10
}