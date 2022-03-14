package utils

import (
	"bytes"
	"encoding/json"
	"errors"
)

// 根据judge的值返回结果,true:ok false:fail
func BoolToResult(judge bool, ok, fail interface{}) interface{} {
	if !judge {
		return fail
	}

	return ok
}

// 将指定map转为具有jsonTag的指针struct
func MapToStructByJson(source map[string]interface{}, target interface{}) (interface{}, error) {
	if len(source) == 0 {
		return nil, errors.New("源map为空")
	}

	buffer := &bytes.Buffer{}
	if err := json.NewEncoder(buffer).Encode(source); err != nil {
		return nil, err
	}

	if err := json.NewDecoder(buffer).Decode(target); err != nil {
		return nil, err
	}

	return target, nil
}

// slice去重
func SliceRemoveDuplicate(arr []interface{}) []interface{} {
	temp := make(map[interface{}]struct{}, len(arr)) // len避免扩容,struct节省空间
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
