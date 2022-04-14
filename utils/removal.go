package utils

type SliceConstraint interface {
	int | int64 | string
}

type R[T SliceConstraint] []T

// slice去除重复数据
func RemoveRespSlice[S SliceConstraint](req []S) []S {
	if len(req) == 0 {
		return nil
	}
	result := make(R[S], 0)
	temp := map[S]struct{}{}

	for _, val := range req {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}

// slice去重
func SliceRemoveDuplicate[S SliceConstraint](arr R[S]) R[S] {
	temp := make(map[S]struct{}, len(arr)) // len避免扩容,struct节省空间
	k := 0

	for _, value := range arr { // 0(n)
		if _, ok := temp[value]; !ok {
			temp[value] = struct{}{}
			arr[k] = value // 记录非重复k,值前移,原地去重 0(n)
			k++
		}
	}

	return arr[:k]
}

// 移除slice中特定的元素
func RemoveSpecificBySlice[S SliceConstraint](arr R[S], param S) R[S] {
	for i := 0; i < len(arr); i++ {
		if arr[i] == param {
			arr = append(arr[:i], arr[i+1:]...) // 无内存分配.提高性能
			i--                                 // 保持正确的索引
		}
	}

	return arr
}
