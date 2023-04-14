// @Author xiaozhaofu 2023/4/14 14:37:00
package E

func InSlice[T int | int32 | int64 | string](items []T, item T) bool {
	m := Slice2Map(items)
	return InMap(m, item)
}
