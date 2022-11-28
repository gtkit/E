// @Author xiaozhaofu 2022/11/28 01:14:00
package hase

// ConvertIntSlice2Map 将字符串 slice 转为 map[string]struct{}。
func IntSlice2Map(sl []int) map[int]struct{} {
	set := make(map[int]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InIntMap(m map[int]struct{}, s int) bool {
	_, ok := m[s]
	return ok
}

func Int64Slice2Map(sl []int64) map[int64]struct{} {
	set := make(map[int64]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InInt64Map(m map[int64]struct{}, s int64) bool {
	_, ok := m[s]
	return ok
}

func Int32Slice2Map(sl []int32) map[int32]struct{} {
	set := make(map[int32]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InInt32Map(m map[int32]struct{}, s int32) bool {
	_, ok := m[s]
	return ok
}

func Int8Slice2Map(sl []int8) map[int8]struct{} {
	set := make(map[int8]struct{}, len(sl))
	for _, v := range sl {
		set[v] = struct{}{}
	}
	return set
}

// InMap 判断字符串是否在 map 中。
func InInt8Map(m map[int8]struct{}, s int8) bool {
	_, ok := m[s]
	return ok
}
