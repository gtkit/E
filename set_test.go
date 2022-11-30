// @Author xiaozhaofu 2022/11/26 20:44:00
package hase

import (
	"testing"
)

var sl = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

// var sl = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func BenchmarkInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InSlice(sl, "m")
		// InSlice(sl, 2)
	}
}

func BenchmarkInMap(b *testing.B) {
	m := Slice2Map(sl)
	for i := 0; i < b.N; i++ {
		InMap(m, "m")
		// InMap(m, 2)
	}
}

// InSlice 判断字符串是否在 slice 中。
func InSlice[T int | int32 | int64 | string](items []T, item T) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
