// @Author xiaozhaofu 2022/11/28 01:07:00
package hase

// ConvertStrSlice2Map 将字符串 slice 转为 map[string]struct{}。
// 使用空结构体 struct{} 作为 value 的类型，因为 struct{} 不占用任何内存空间
// fmt.Println(unsafe.Sizeof(bool(false))) // 1
// fmt.Println(unsafe.Sizeof(struct{}{}))  // 0
func StrSlice2Map(sl []string) map[string]struct{} {
	set := make(map[string]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InStrMap(m map[string]struct{}, s string) bool {
	_, ok := m[s]
	return ok
}
