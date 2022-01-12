package duplicate

// slice去除重复数据
func RemoveRespSlice(req []int64) []int64 {
	if len(req) == 0 {
		return nil
	}
	result := make([]int64, 0)
	temp := map[int64]struct{}{}

	for _, val := range req {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}
