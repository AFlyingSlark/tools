package utils

import (
	"bytes"
	"encoding/json"
	"errors"
)

type BoolToResultType interface {
	int | int64 | string | []int64 | []string | func()
}

// 根据judge的值返回结果,true:ok false:fail
func BoolToResult[T BoolToResultType](judge bool, ok, fail T) T {
	if !judge {
		return fail
	}

	return ok
}

// 将指定map转为具有jsonTag的指针struct
func MapToStructByJson(source map[string]any, target any) (any, error) {
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

// slice交集
func SliceIntersection[T int64 | string](first []T, last []T) []T {
	result := make([]T, 0)
	temp := make(map[T]struct{})

	for _, value := range first {
		if _, ok := temp[value]; !ok {
			temp[value] = struct{}{}
		}
	}

	for _, val := range last {
		if _, ok := temp[val]; ok {
			result = append(result, val)
		}
	}

	return result
}
