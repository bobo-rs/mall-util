package util

import (
	"github.com/gogf/gf/v2/text/gstr"
	"regexp"
	"strings"
)

// SplitWords 字符串分拆关键词
// Example
// s := "264827482974
//
//	57839579357
//	73973457895397"
//
// SplitWords(s)
// Output
// []string{"264827482974", "57839579357", "73973457895397"}
func SplitWords(s string, compiles ...string) []string {
	if len(s) == 0 {
		return []string{}
	}
	// 匹配条件
	compile := `[\n\r|\t|,，]+`
	if len(compiles) > 0 {
		compile = compiles[0]
	}
	// 匹配
	regex := regexp.MustCompile(compile)
	words := regex.Split(s, -1)
	if len(words) == 0 {
		return words
	}
	// 对数据进行过滤
	for i, v := range words {
		words[i] = Filter(v)
	}
	return words
}

// Filter 过滤字符串空格和转义字符
func Filter(s string) string {
	if len(s) == 0 {
		return ""
	}
	return gstr.AddSlashes(strings.TrimSpace(s))
}
