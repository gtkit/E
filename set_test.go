// @Author xiaozhaofu 2022/11/26 20:44:00
package hase

import (
	"testing"
)

var sl = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func BenchmarkInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InSlice(sl, "m")
	}
}

func BenchmarkInMap(b *testing.B) {
	m := StrSlice2Map(sl)
	for i := 0; i < b.N; i++ {
		InStrMap(m, "m")
	}
}

// InSlice 判断字符串是否在 slice 中。
func InSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
