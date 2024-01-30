package util

import (
	"math"
)

// Add 加法,两数及以上的数值累计相加之和，并保留小数点后两位
// Example:
// Add(100, 2000, 2, 100.01111)
// Output: 2200.01
func Add(x, y interface{}, decimal float64, z ...interface{}) float64 {
	sum := Float64(x) + Float64(y)
	if len(z) > 0 {
		for _, v := range z {
			sum += Float64(v)
		}
	}
	return Scale(sum, decimal)
}

// Sub 减法,两数及以上的数值累计之差，并保留小数点后两位
// Example:
// Sub(100, 2000, 2, 100.01111)
// Output: -2000.01
func Sub(x, y interface{}, decimal float64, z ...interface{}) float64 {
	s := Float64(x) - Float64(y)
	if len(z) > 0 {
		for _, v := range z {
			s -= Float64(v)
		}
	}
	return Scale(s, decimal)
}

// Mul 乘法,两数及以上的数值乘积，并保留小数点后两位
// Example:
// Mul(100, 2000, 2, 100.01111)
// Output: 2.0002222e+07
func Mul(x, y interface{}, decimal float64, z ...interface{}) float64 {
	s := Float64(x) * Float64(y)
	if len(z) > 0 {
		for _, v := range z {
			s *= Float64(v)
		}
	}
	if s == 0 {
		return 0
	}
	return Scale(s, decimal)
}

// Div 除法， 两数相除
// Example:
// Sub(999, 8989.5645, 2)
// Output: 0.11
func Div(x, y interface{}, decimal float64) float64 {
	if Float64(y) == 0 || Float64(x) == 0 {
		return 0
	}
	s := Float64(x) / Float64(y)
	return Scale(s, decimal)
}

// Scale 保留小数点-四舍五入算法
// Example:
// Scale(98.88888, 2) 保留两位小数
// Output: 98.89
func Scale(s, decimal float64) float64 {
	// 保留小数点，注：避免小数点小于等于0，计算错误
	var d float64 = 1
	if decimal > 1 {
		d = math.Pow(10, decimal) // 10的N次方
	}
	return math.Round(s*d) / d
}
