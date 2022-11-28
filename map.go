// @Author xiaozhaofu 2022/11/28 01:07:00
package hase

import (
	"fmt"
)

// Slice2Map 将字符串 slice 转为 map[T]struct{}。
// 使用空结构体 struct{} 作为 value 的类型，因为 struct{} 不占用任何内存空间
// fmt.Println(unsafe.Sizeof(bool(false))) // 1
// fmt.Println(unsafe.Sizeof(struct{}{}))  // 0
func Slice2Map[T int | int32 | int64 | string](sl []T) map[T]struct{} {
	fmt.Printf("sl type name----%T\n", sl)
	set := make(map[T]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InMap[T int | int32 | int64 | string](m map[T]struct{}, s T) bool {
	_, ok := m[s]
	return ok
}
