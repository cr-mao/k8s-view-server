package utils

import (
	"strconv"
	"time"
)

func StringSlice2IntSlice(strArr []string) []int {
	res := make([]int, len(strArr))

	for index, val := range strArr {
		res[index], _ = strconv.Atoi(val)
	}

	return res
}

func StringSlice2Int32Slice(strArr []string) []int32 {
	res := make([]int32, len(strArr))

	for index, val := range strArr {
		id, _ := strconv.Atoi(val)
		res[index] = int32(id)
	}

	return res
}

func Int32SliceRemove(arrBig []int32, arrSmall []int32) []int32 {
	for _, v := range arrSmall {
		//arrBig = Int32SliceRemoveOne(arrBig, v)
		for i, bigItem := range arrBig {
			if bigItem == v {
				///验证过，没问题，但是 arrBig[i+1:]用法中 大于len(arr)的时候会有问题
				//arrBig[:i] i为0时候没问题，i大于len(arr)的时候会有问题
				arrBig = append(arrBig[:i], arrBig[i+1:]...)
				break
			}
		}
	}
	return arrBig
}

// /数组中某个值的元素都移除掉
func Int32SliceRemoveValue(arr []int32, elem int32) []int32 {
	r := arr[:0]
	for _, v := range arr {
		if v != elem {
			r = append(r, v)
		}
	}
	return r
}

// 只移除一个值为elem的元素，首个
func Int32SliceRemoveOne(arr []int32, elem int32) []int32 {
	for i, bigItem := range arr {
		if bigItem == elem {
			///验证过，没问题，但是 arrBig[i+1:]用法中 大于len(arr)的时候会有问题
			//arrBig[:i] i为0时候没问题，i大于len(arr)的时候会有问题
			ret := append(arr[:i], arr[i+1:]...)
			return ret
		}
	}
	return arr
}

func Int32SliceContains(arr []int32, val int32) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
func StringSliceContains(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func GetTimestamp() int64 {
	now := time.Now().Unix()
	return now
}

func FixInt32(value int32, min, max int32) int32 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}
func FixFloat64(value float64, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// 数组翻转
func RevertIntSlice(arr []int32) {
	slen := len(arr)
	for i := 0; i < slen/2; i++ {
		j := slen - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func MapContainsKey(m map[int32]int64, key int32) bool {
	if m == nil {
		return false
	}
	if _, ok := m[key]; ok {
		return ok
	}
	return false
}
func MapAdd(m map[int32]int64, key int32, v int64) int64 {
	if m == nil {
		return v
	}
	exist := int64(0)
	if cnt, ok := m[key]; ok {
		exist = cnt
	}
	exist += v
	m[key] = exist
	return exist
}
func MapMerge(m1, m2 map[int32]int64) map[int32]int64 {
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}
	m := map[int32]int64{}
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}
func MapMergeInt32(m1, m2 map[int32]int32) map[int32]int32 {
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}
	m := map[int32]int32{}
	for k, v := range m1 {
		m[k] = v
	}
	for k, v := range m2 {
		m[k] = v
	}
	return m
}
